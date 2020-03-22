/*
 * Copyright 2015, gRPC Authors All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.grpc.internal;

import static io.grpc.ConnectivityState.CONNECTING;
import static io.grpc.ConnectivityState.IDLE;
import static io.grpc.ConnectivityState.READY;
import static io.grpc.ConnectivityState.SHUTDOWN;
import static io.grpc.ConnectivityState.TRANSIENT_FAILURE;

import com.google.common.annotations.VisibleForTesting;
import com.google.common.base.Preconditions;
import com.google.common.base.Stopwatch;
import com.google.common.base.Supplier;
import com.google.errorprone.annotations.ForOverride;
import io.grpc.ConnectivityState;
import io.grpc.ConnectivityStateInfo;
import io.grpc.EquivalentAddressGroup;
import io.grpc.Status;
import java.net.SocketAddress;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.annotation.Nullable;
import javax.annotation.concurrent.GuardedBy;
import javax.annotation.concurrent.ThreadSafe;

/**
 * Transports for a single {@link SocketAddress}.
 */
@ThreadSafe
final class InternalSubchannel implements WithLogId {
  private static final Logger log = Logger.getLogger(InternalSubchannel.class.getName());

  private final LogId logId = LogId.allocate(getClass().getName());
  private final String authority;
  private final String userAgent;
  private final BackoffPolicy.Provider backoffPolicyProvider;
  private final Callback callback;
  private final ClientTransportFactory transportFactory;
  private final ScheduledExecutorService scheduledExecutor;

  // File-specific convention: methods without GuardedBy("lock") MUST NOT be called under the lock.
  private final Object lock = new Object();

  // File-specific convention:
  //
  // 1. In a method without GuardedBy("lock"), executeLater() MUST be followed by a drain() later in
  // the same method.
  //
  // 2. drain() MUST NOT be called under "lock".
  //
  // 3. Every synchronized("lock") must be inside a try-finally which calls drain() in "finally".
  private final ChannelExecutor channelExecutor;

  @GuardedBy("lock")
  private EquivalentAddressGroup addressGroup;

  /**
   * The index of the address corresponding to pendingTransport/activeTransport, or 0 if both are
   * null.
   */
  @GuardedBy("lock")
  private int addressIndex;

  /**
   * The policy to control back off between reconnects. Non-{@code null} when a reconnect task is
   * scheduled.
   */
  @GuardedBy("lock")
  private BackoffPolicy reconnectPolicy;

  /**
   * Timer monitoring duration since entering CONNECTING state.
   */
  @GuardedBy("lock")
  private final Stopwatch connectingTimer;

  @GuardedBy("lock")
  @Nullable
  private ScheduledFuture<?> reconnectTask;

  @GuardedBy("lock")
  private boolean reconnectCanceled;

  /**
   * All transports that are not terminated. At the very least the value of {@link #activeTransport}
   * will be present, but previously used transports that still have streams or are stopping may
   * also be present.
   */
  @GuardedBy("lock")
  private final Collection<ConnectionClientTransport> transports =
      new ArrayList<ConnectionClientTransport>();

  // Must only be used from channelExecutor
  private final InUseStateAggregator<ConnectionClientTransport> inUseStateAggregator =
      new InUseStateAggregator<ConnectionClientTransport>() {
        @Override
        void handleInUse() {
          callback.onInUse(InternalSubchannel.this);
        }

        @Override
        void handleNotInUse() {
          callback.onNotInUse(InternalSubchannel.this);
        }
      };

  /**
   * The to-be active transport, which is not ready yet.
   */
  @GuardedBy("lock")
  @Nullable
  private ConnectionClientTransport pendingTransport;

  /**
   * The transport for new outgoing requests. 'lock' must be held when assigning to it. Non-null
   * only in READY state.
   */
  @Nullable
  private volatile ManagedClientTransport activeTransport;

  @GuardedBy("lock")
  private ConnectivityStateInfo state = ConnectivityStateInfo.forNonError(IDLE);

  private final ProxyDetector proxyDetector;

  @GuardedBy("lock")
  private Status shutdownReason;


  InternalSubchannel(EquivalentAddressGroup addressGroup, String authority, String userAgent,
      BackoffPolicy.Provider backoffPolicyProvider,
      ClientTransportFactory transportFactory, ScheduledExecutorService scheduledExecutor,
      Supplier<Stopwatch> stopwatchSupplier, ChannelExecutor channelExecutor, Callback callback,
      ProxyDetector proxyDetector) {
    this.addressGroup = Preconditions.checkNotNull(addressGroup, "addressGroup");
    this.authority = authority;
    this.userAgent = userAgent;
    this.backoffPolicyProvider = backoffPolicyProvider;
    this.transportFactory = transportFactory;
    this.scheduledExecutor = scheduledExecutor;
    this.connectingTimer = stopwatchSupplier.get();
    this.channelExecutor = channelExecutor;
    this.callback = callback;
    this.proxyDetector = proxyDetector;
  }

  /**
   * Returns a READY transport that will be used to create new streams.
   *
   * <p>Returns {@code null} if the state is not READY.  Will try to connect if state is IDLE.
   */
  @Nullable
  ClientTransport obtainActiveTransport() {
    ClientTransport savedTransport = activeTransport;
    if (savedTransport != null) {
      return savedTransport;
    }
    try {
      synchronized (lock) {
        savedTransport = activeTransport;
        // Check again, since it could have changed before acquiring the lock
        if (savedTransport != null) {
          return savedTransport;
        }
        if (state.getState() == IDLE) {
          gotoNonErrorState(CONNECTING);
          startNewTransport();
        }
      }
    } finally {
      channelExecutor.drain();
    }
    return null;
  }

  @GuardedBy("lock")
  private void startNewTransport() {
    Preconditions.checkState(reconnectTask == null, "Should have no reconnectTask scheduled");

    if (addressIndex == 0) {
      connectingTimer.reset().start();
    }
    List<SocketAddress> addrs = addressGroup.getAddresses();
    final SocketAddress address = addrs.get(addressIndex);

    ProxyParameters proxy = proxyDetector.proxyFor(address);
    ConnectionClientTransport transport =
        transportFactory.newClientTransport(address, authority, userAgent, proxy);
    if (log.isLoggable(Level.FINE)) {
      log.log(Level.FINE, "[{0}] Created {1} for {2}",
          new Object[] {logId, transport.getLogId(), address});
    }
    pendingTransport = transport;
    transports.add(transport);
    Runnable runnable = transport.start(new TransportListener(transport, address));
    if (runnable != null) {
      channelExecutor.executeLater(runnable);
    }
  }

  /**
   * Only called after all addresses attempted and failed (TRANSIENT_FAILURE).
   * @param status the causal status when the channel begins transition to
   *     TRANSIENT_FAILURE.
   */
  @GuardedBy("lock")
  private void scheduleBackoff(final Status status) {
    class EndOfCurrentBackoff implements Runnable {
      @Override
      public void run() {
        try {
          synchronized (lock) {
            reconnectTask = null;
            if (reconnectCanceled) {
              // Even though cancelReconnectTask() will cancel this task, the task may have already
              // started when it's being canceled.
              return;
            }
            gotoNonErrorState(CONNECTING);
            startNewTransport();
          }
        } catch (Throwable t) {
          log.log(Level.WARNING, "Exception handling end of backoff", t);
        } finally {
          channelExecutor.drain();
        }
      }
    }

    gotoState(ConnectivityStateInfo.forTransientFailure(status));
    if (reconnectPolicy == null) {
      reconnectPolicy = backoffPolicyProvider.get();
    }
    long delayNanos =
        reconnectPolicy.nextBackoffNanos() - connectingTimer.elapsed(TimeUnit.NANOSECONDS);
    if (log.isLoggable(Level.FINE)) {
      log.log(Level.FINE, "[{0}] Scheduling backoff for {1} ns", new Object[]{logId, delayNanos});
    }
    Preconditions.checkState(reconnectTask == null, "previous reconnectTask is not done");
    reconnectCanceled = false;
    reconnectTask = scheduledExecutor.schedule(
        new LogExceptionRunnable(new EndOfCurrentBackoff()),
        delayNanos,
        TimeUnit.NANOSECONDS);
  }

  /**
   * Immediately attempt to reconnect if the current state is TRANSIENT_FAILURE. Otherwise this
   * method has no effect.
   */
  void resetConnectBackoff() {
    try {
      synchronized (lock) {
        if (state.getState() != TRANSIENT_FAILURE) {
          return;
        }
        cancelReconnectTask();
        gotoNonErrorState(CONNECTING);
        startNewTransport();
      }
    } finally {
      channelExecutor.drain();
    }
  }

  @GuardedBy("lock")
  private void gotoNonErrorState(ConnectivityState newState) {
    gotoState(ConnectivityStateInfo.forNonError(newState));
  }

  @GuardedBy("lock")
  private void gotoState(final ConnectivityStateInfo newState) {
    if (state.getState() != newState.getState()) {
      Preconditions.checkState(state.getState() != SHUTDOWN,
          "Cannot transition out of SHUTDOWN to " + newState);
      state = newState;
      channelExecutor.executeLater(new Runnable() {
          @Override
          public void run() {
            callback.onStateChange(InternalSubchannel.this, newState);
          }
        });
    }
  }

  /** Replaces the existing addresses, avoiding unnecessary reconnects. */
  public void updateAddresses(EquivalentAddressGroup newAddressGroup) {
    ManagedClientTransport savedTransport = null;
    try {
      synchronized (lock) {
        EquivalentAddressGroup oldAddressGroup = addressGroup;
        addressGroup = newAddressGroup;
        if (state.getState() == READY || state.getState() == CONNECTING) {
          SocketAddress address = oldAddressGroup.getAddresses().get(addressIndex);
          int newIndex = newAddressGroup.getAddresses().indexOf(address);
          if (newIndex != -1) {
            addressIndex = newIndex;
          } else {
            // Forced to drop the connection
            if (state.getState() == READY) {
              savedTransport = activeTransport;
              activeTransport = null;
              addressIndex = 0;
              gotoNonErrorState(IDLE);
            } else {
              savedTransport = pendingTransport;
              pendingTransport = null;
              addressIndex = 0;
              startNewTransport();
            }
          }
        }
      }
    } finally {
      channelExecutor.drain();
    }
    if (savedTransport != null) {
      savedTransport.shutdown(
          Status.UNAVAILABLE.withDescription(
              "InternalSubchannel closed transport due to address change"));
    }
  }

  public void shutdown(Status reason) {
    ManagedClientTransport savedActiveTransport;
    ConnectionClientTransport savedPendingTransport;
    try {
      synchronized (lock) {
        if (state.getState() == SHUTDOWN) {
          return;
        }
        shutdownReason = reason;
        gotoNonErrorState(SHUTDOWN);
        savedActiveTransport = activeTransport;
        savedPendingTransport = pendingTransport;
        activeTransport = null;
        pendingTransport = null;
        addressIndex = 0;
        if (transports.isEmpty()) {
          handleTermination();
          if (log.isLoggable(Level.FINE)) {
            log.log(Level.FINE, "[{0}] Terminated in shutdown()", logId);
          }
        }  // else: the callback will be run once all transports have been terminated
        cancelReconnectTask();
      }
    } finally {
      channelExecutor.drain();
    }
    if (savedActiveTransport != null) {
      savedActiveTransport.shutdown(reason);
    }
    if (savedPendingTransport != null) {
      savedPendingTransport.shutdown(reason);
    }
  }

  @GuardedBy("lock")
  private void handleTermination() {
    channelExecutor.executeLater(new Runnable() {
        @Override
        public void run() {
          callback.onTerminated(InternalSubchannel.this);
        }
      });
  }

  private void handleTransportInUseState(
      final ConnectionClientTransport transport, final boolean inUse) {
    channelExecutor.executeLater(new Runnable() {
        @Override
        public void run() {
          inUseStateAggregator.updateObjectInUse(transport, inUse);
        }
      }).drain();
  }

  void shutdownNow(Status reason) {
    shutdown(reason);
    Collection<ManagedClientTransport> transportsCopy;
    try {
      synchronized (lock) {
        transportsCopy = new ArrayList<ManagedClientTransport>(transports);
      }
    } finally {
      channelExecutor.drain();
    }
    for (ManagedClientTransport transport : transportsCopy) {
      transport.shutdownNow(reason);
    }
  }

  EquivalentAddressGroup getAddressGroup() {
    try {
      synchronized (lock) {
        return addressGroup;
      }
    } finally {
      channelExecutor.drain();
    }
  }

  @GuardedBy("lock")
  private void cancelReconnectTask() {
    if (reconnectTask != null) {
      reconnectTask.cancel(false);
      reconnectCanceled = true;
      reconnectTask = null;
      reconnectPolicy = null;
    }
  }

  @Override
  public LogId getLogId() {
    return logId;
  }

  @VisibleForTesting
  ConnectivityState getState() {
    try {
      synchronized (lock) {
        return state.getState();
      }
    } finally {
      channelExecutor.drain();
    }
  }

  /** Listener for real transports. */
  private class TransportListener implements ManagedClientTransport.Listener {
    final ConnectionClientTransport transport;
    final SocketAddress address;

    TransportListener(ConnectionClientTransport transport, SocketAddress address) {
      this.transport = transport;
      this.address = address;
    }

    @Override
    public void transportReady() {
      if (log.isLoggable(Level.FINE)) {
        log.log(Level.FINE, "[{0}] {1} for {2} is ready",
            new Object[] {logId, transport.getLogId(), address});
      }
      Status savedShutdownReason;
      try {
        synchronized (lock) {
          savedShutdownReason = shutdownReason;
          reconnectPolicy = null;
          if (savedShutdownReason != null) {
            // activeTransport should have already been set to null by shutdown(). We keep it null.
            Preconditions.checkState(activeTransport == null,
                "Unexpected non-null activeTransport");
          } else if (pendingTransport == transport) {
            gotoNonErrorState(READY);
            activeTransport = transport;
            pendingTransport = null;
          }
        }
      } finally {
        channelExecutor.drain();
      }
      if (savedShutdownReason != null) {
        transport.shutdown(savedShutdownReason);
      }
    }

    @Override
    public void transportInUse(boolean inUse) {
      handleTransportInUseState(transport, inUse);
    }

    @Override
    public void transportShutdown(Status s) {
      if (log.isLoggable(Level.FINE)) {
        log.log(Level.FINE, "[{0}] {1} for {2} is being shutdown with status {3}",
            new Object[] {logId, transport.getLogId(), address, s});
      }
      try {
        synchronized (lock) {
          if (state.getState() == SHUTDOWN) {
            return;
          }
          if (activeTransport == transport) {
            gotoNonErrorState(IDLE);
            activeTransport = null;
            addressIndex = 0;
          } else if (pendingTransport == transport) {
            Preconditions.checkState(state.getState() == CONNECTING,
                "Expected state is CONNECTING, actual state is %s", state.getState());
            addressIndex++;
            // Continue reconnect if there are still addresses to try.
            if (addressIndex >= addressGroup.getAddresses().size()) {
              pendingTransport = null;
              addressIndex = 0;
              // Initiate backoff
              // Transition to TRANSIENT_FAILURE
              scheduleBackoff(s);
            } else {
              startNewTransport();
            }
          }
        }
      } finally {
        channelExecutor.drain();
      }
    }

    @Override
    public void transportTerminated() {
      if (log.isLoggable(Level.FINE)) {
        log.log(Level.FINE, "[{0}] {1} for {2} is terminated",
            new Object[] {logId, transport.getLogId(), address});
      }
      handleTransportInUseState(transport, false);
      try {
        synchronized (lock) {
          transports.remove(transport);
          if (state.getState() == SHUTDOWN && transports.isEmpty()) {
            if (log.isLoggable(Level.FINE)) {
              log.log(Level.FINE, "[{0}] Terminated in transportTerminated()", logId);
            }
            handleTermination();
          }
        }
      } finally {
        channelExecutor.drain();
      }
      Preconditions.checkState(activeTransport != transport,
          "activeTransport still points to this transport. "
          + "Seems transportShutdown() was not called.");
    }
  }

  // All methods are called in channelExecutor, which is a serializing executor.
  abstract static class Callback {
    /**
     * Called when the subchannel is terminated, which means it's shut down and all transports
     * have been terminated.
     */
    @ForOverride
    void onTerminated(InternalSubchannel is) { }

    /**
     * Called when the subchannel's connectivity state has changed.
     */
    @ForOverride
    void onStateChange(InternalSubchannel is, ConnectivityStateInfo newState) { }

    /**
     * Called when the subchannel's in-use state has changed to true, which means at least one
     * transport is in use.
     */
    @ForOverride
    void onInUse(InternalSubchannel is) { }

    /**
     * Called when the subchannel's in-use state has changed to false, which means no transport is
     * in use.
     */
    @ForOverride
    void onNotInUse(InternalSubchannel is) { }
  }
}
