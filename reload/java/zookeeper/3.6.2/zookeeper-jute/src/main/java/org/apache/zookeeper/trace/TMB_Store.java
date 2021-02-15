package org.apache.zookeeper.trace;

import org.apache.jute.Record;
import org.apache.zookeeper.proto.NullPointerResponse;
import org.apache.zookeeper.server.quorum.QuorumPacket;

import java.io.ByteArrayInputStream;
import java.util.*;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class TMB_Store {
    private static TMB_Store instance = null;

    public static int REQUESTER = 1;
    public static int RESPONSER = 2;

    public static TMB_Store getInstance() {
        if (instance == null) {
            instance = new TMB_Store();

            instance.clientPluginTraces = new HashMap<>();
            instance.quorumTraces = new HashMap<>();
            instance.quorumLock = new ReentrantReadWriteLock();
        }

        return instance;
    }

    class ClientPluginTrace {
        private TMB_Trace trace;

        ClientPluginTrace(TMB_Trace trace) {
            this.trace = trace;
        }
    }

    class QuorumTrace {
        private String quorumIdStr;
        public int recorder;

        private TMB_Trace trace;

        QuorumTrace(long quorumId, TMB_Trace trace) { // make sure trace has events
            this.quorumIdStr = String.format("quorum-%d", quorumId);
            this.trace = trace;

            if (trace.getEvents().get(0).getService().equals(this.quorumIdStr)) {
                this.recorder = REQUESTER;
            } else {
                this.recorder = RESPONSER;
            }
        }
    }

    public class QuorumTraces {
        private long quorumId;
        private Map<Long, QuorumTrace> traces; // Map<traceId, Trace>
        private ReentrantReadWriteLock lock;

        QuorumTraces(long quorumId) {
            this.quorumId = quorumId;
            this.traces = new HashMap<>();
            this.lock = new ReentrantReadWriteLock();
        }

        public void printAllJSON() {
            lock.readLock().lock();
            for (Long key: traces.keySet()) {
                TMB_Helper.printf("[TMB_Store] [quorum-%d] %d: %d, %s\n", quorumId, key, traces.get(key).recorder, traces.get(key).trace.toJSON());
            }
            lock.readLock().unlock();
        }

        public void setTrace(TMB_Trace trace_) {
            lock.writeLock().lock();
            QuorumTrace quorumTrace = traces.get(trace_.getId());
            if (quorumTrace != null) {
                mergeEvents(quorumTrace.trace, trace_.getEvents());
            } else {
                quorumTrace = new QuorumTrace(quorumId, trace_.copy());
            }
            traces.put(trace_.getId(), quorumTrace);
            lock.writeLock().unlock();
        }

        public QuorumTrace getQuorumTrace(long traceId) {
            lock.readLock().lock();
            QuorumTrace quorumTrace = traces.get(traceId);
            lock.readLock().unlock();

            return quorumTrace;
        }

        public QuorumTrace removeQuorumTrace(long traceId) {
            lock.readLock().lock();
            QuorumTrace quorumTrace = traces.remove(traceId);
            lock.readLock().unlock();

            return quorumTrace;
        }
    }

    private Map<Long, ClientPluginTrace> clientPluginTraces; // Map<traceId, TMB_Trace>
    private Map<Long, QuorumTraces> quorumTraces; // Map<quorumId, QuorumTrace>
    private ReentrantReadWriteLock quorumLock;

    public void quorumSendRequest(long quorumId, TMB_Trace trace_) {
        quorumSetTrace(quorumId, trace_);
    }

    public void quorumReceiveRequest(long quorumId, TMB_Trace trace_) {
        quorumSetTrace(quorumId, trace_);
    }

    public void quorumReceiveResponse(long quorumId, TMB_Trace trace_) {
        quorumSetTrace(quorumId, trace_);
    }

    public void quorumSendResponse(long quorumId, TMB_Trace trace_) {
        quorumSetTrace(quorumId, trace_);
    }

    public void quorumQuit(long quorumId, TMB_Trace trace_) {
        quorumLock.readLock().lock();
        QuorumTraces quorumTraces = getQuorumTraces(quorumId);
        quorumLock.readLock().unlock();

        QuorumTrace quorumTrace = quorumTraces.getQuorumTrace(trace_.getId());
        if (quorumTrace != null && quorumTrace.recorder == RESPONSER) {
            quorumTraces.removeQuorumTrace(trace_.getId());
        }
    }

    public void quorumSetTrace(long quorumId, TMB_Trace trace_) {
        QuorumTraces quorumTraces = getQuorumTraces(quorumId);
        quorumTraces.setTrace(trace_);
    }

    public TMB_Trace quorumGetTrace(long quorumId, long traceId) { // unsafe
        QuorumTraces quorumTraces = getQuorumTraces(quorumId);

        quorumTraces.lock.readLock().lock();
        QuorumTrace quorumTrace = quorumTraces.getQuorumTrace(traceId);
        quorumTraces.lock.readLock().unlock();

        return quorumTrace.trace;
    }

    public QuorumTraces getQuorumTraces(long quorumId) {
        quorumLock.readLock().lock();
        QuorumTraces quorumTraces = this.quorumTraces.get(quorumId);
        quorumLock.readLock().unlock();

        if (quorumTraces == null) {
            quorumLock.writeLock().lock();
            quorumTraces = this.quorumTraces.get(quorumId);
            if (quorumTraces == null) {
                quorumTraces = new QuorumTraces(quorumId);
                this.quorumTraces.put(quorumId, quorumTraces);
            }
            quorumLock.writeLock().unlock();
        }

        return quorumTraces;
    }

    public static void mergeEvents(TMB_Trace trace, List<TMB_Event> events_) {
        List<TMB_Event> events = trace.getEvents();
        for (TMB_Event event_: events_) {
            boolean found = false;
            for (TMB_Event event: events) {
                found = event.equals(event_);
                if (found) {
                    break;
                }
            }
            if (!found) {
                events.add(event_);
            }
        }
        trace.setEvents(events);
    }

    // called when initializing TMB_ClientPlugin
    public void setClientPluginTrace(TMB_Trace trace) {
        if (trace == null) {
            return;
        }
        long traceId = trace.getId();
        if (traceId == 0) {
            return;
        }

        lock.writeLock().lock();
        if (!clientPluginTraces.containsKey(traceId)) {
            clientPluginTraces.put(traceId, new ClientPluginTrace(trace));
        }
        lock.writeLock().unlock();
    }

    public void updateClientPluginTrace(TMB_Trace trace) {
        if (trace == null) {
            return;
        }
        long traceId = trace.getId();
        if (traceId == 0) {
            return;
        }

        lock.writeLock().lock();
        ClientPluginTrace pluginTrace = clientPluginTraces.get(traceId);
        if (pluginTrace != null) {

        } else {
            TMB_Helper.printf("unreachable code, trace: %s", trace.toJSON());
        }

        lock.writeLock().unlock();
    }

    private static Map<Long, TMB_Trace> thread_traces = new HashMap<>();
    private static Map<Long, TMB_Trace> server_traces = new HashMap<>();
    private static ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    public static void clearServerTraces() {
        lock.writeLock().lock();
        server_traces.clear();
        lock.writeLock().unlock();
    }

    public static void serverSetTrace(long serverId, TMB_Trace trace) {
        if (trace.getId() == 0) {
            return;
        }

        lock.writeLock().lock();
        server_traces.put(serverId, trace);
        lock.writeLock().unlock();
    }

    public static void serverMergeTrace(long serverId, TMB_Trace trace) {
        if (trace == null || trace.getId() == 0) {
            return;
        }

        List<TMB_Event> events = trace.getEvents();
        int reqEvent = (int) trace.getReqEvent();
        if (events.size() <= reqEvent) {
            return;
        }

        events = events.subList(reqEvent, events.size());

        lock.writeLock().lock();
        TMB_Trace trace_ = server_traces.get(serverId);
        if (trace_ == null || trace_.getId() != trace.getId()) {
            lock.writeLock().unlock();

            return;
        }

        List<TMB_Event> events_ = trace_.getEvents();
        for (TMB_Event event: events) {
            boolean found = false;
            for (TMB_Event event_: events_) {
                found = event_.equals(event);
                if (found) {
                    break;
                }
            }
            if (!found) {
                events_.add(event);
            }
        }
        trace_.setEvents(events_);

        lock.writeLock().unlock();
    }

    public static TMB_Trace serverGetTrace(long serverId) {
//        lock.readLock().lock();
//        TMB_Trace trace = server_traces.get(serverId);
//        lock.readLock().unlock();

        return null;
    }

    public static void callerAppendEventsByThreadIdUnsafe(long threadId, TMB_Trace trace) {
        TMB_Trace trace_ = thread_traces.get(threadId);

        if (trace_ == null) {
            thread_traces.put(threadId, trace);
        } else {
//            TMB_Helper.println("before appending, length: " + trace_.getEvents().size() + " + " + trace.getEvents().size());

            List<TMB_Event> events_ = trace_.getEvents();
            List<TMB_Event> events = trace.getEvents();

            events_.addAll(events);
            trace_.setEvents(events_);

//            TMB_Helper.println("after appending, length: " + trace_.getEvents().size());
        }
    }

    public static TMB_Trace getByThreadId(long threadId) {
        lock.readLock().lock();
        TMB_Trace trace = thread_traces.get(threadId);
        lock.readLock().unlock();

        return trace;
    }

    // TODO: 3MileBeach temp function
    public static Map<Long, TMB_Trace> getAllThreads() {
        Map<Long, TMB_Trace> result = new HashMap<>();
        lock.readLock().lock();
        thread_traces.forEach(result::put);
        lock.readLock().unlock();

        return result;
    }

    public static void removeByThreadId(long threadId) {
        lock.writeLock().lock();
        thread_traces.remove(threadId);
        lock.writeLock().unlock();
    }

    /**
     * Server receives request from client or upstream server
     * RequestProcessor -> calleeInbound -> process request -> calleeOutbound -> send response
     */
    // to be called after ByteBufferInputStream.byteBuffer2Record(buffer, record)
    public static void calleeInbound(int serviceId, String service, Record request) {
        TMB_Trace trace = request.getTrace();
        if (trace == null || trace.getId() == 0) {
            TMB_Helper.printf(3, "[%s] callee inbound receives request with empty trace: %s, (%s)\n", service, TMB_Helper.getClassName(request), TMB_Helper.getString(request));
            return;
        }

        List<TMB_Event> events = trace.getEvents();
        // TODO: 3MileBeach
//        if (events.size() != 1) {
//            TMB_Helper.printf("[%s] callee inbound receives request with invalid trace: %s, (%s)\n", service, TMB_Helper.getClassName(request), TMB_Helper.getString(request));
//
//            return;
//        }

        TMB_Helper.printf(3, "[%s] callee inbound receives request: %s, (%s)\n", service, TMB_Helper.getClassName(request), TMB_Helper.getString(request));

        long threadId = Thread.currentThread().getId();
        TMB_Event preEvent = events.get(0);
        TMB_Event event = new TMB_Event(TMB_Event.RECORD_RECV, TMB_Helper.currentTimeNanos(), preEvent.getMessage_name(), preEvent.getUuid(), service);
        events.set(0, event);
        trace.setEvents(events);

        lock.writeLock().lock();
        thread_traces.put(threadId, trace);
        lock.writeLock().unlock();

        getInstance().quorumSetTrace(serviceId, trace);
    }

    public static void calleeOutbound(int serviceId, String service, Record response) {
        long threadId = Thread.currentThread().getId();

        lock.writeLock().lock();
        TMB_Trace trace = thread_traces.get(threadId);
        thread_traces.remove(threadId);
        lock.writeLock().unlock();

        TMB_Trace trace_ = getInstance().quorumGetTrace(serviceId, trace.getId());

        if (trace == null) {
            TMB_Helper.printf(3, "[%s] callee outbound ejects response without trace: %s, (%s)\n", service, TMB_Helper.getClassName(response), TMB_Helper.getString(response));

            return;
        }

        TMB_Event preEvent = trace.getEvents().get(0);
        mergeEvents(trace, trace_.getEvents());

        TMB_Event event = new TMB_Event(TMB_Event.RECORD_SEND, TMB_Helper.currentTimeNanos(), TMB_Helper.getClassName(response), preEvent.getUuid(), service);

        trace.addEvent(event);
        response.setTrace(trace);

        TMB_Helper.printf(3, "[%s] callee outbound ejects response: %s, (%s)\n", service, TMB_Helper.getClassName(response), TMB_Helper.getString(response));
    }
}
