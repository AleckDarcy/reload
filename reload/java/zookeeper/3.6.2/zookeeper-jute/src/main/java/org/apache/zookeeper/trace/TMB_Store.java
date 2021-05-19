package org.apache.zookeeper.trace;

import org.apache.jute.Record;

import java.util.*;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class TMB_Store {
    private static TMB_Store instance = new TMB_Store();

    public static int REQUESTER = 1;
    public static int RESPONDER = 2;

    private TMB_Store() {
        this.clientPluginTraces = new HashMap<>();
        this.quorumTraces = new HashMap<>();
        this.quorumLock = new ReentrantReadWriteLock();
    }

    public static TMB_Store getInstance() {
        return instance;
    }

    class ClientPluginTrace {
        private TMB_Trace trace;

        ClientPluginTrace(TMB_Trace trace) {
            this.trace = trace;
        }
    }

    public static class ProcessorMeta {
        private QuorumMeta quorumMeta;
        private String name;

        public ProcessorMeta(QuorumMeta quorumMeta, Object processor) {
            this.quorumMeta = quorumMeta;
            this.name = TMB_Helper.getClassName(processor.getClass());
        }

        public QuorumMeta getQuorumMeta() {
            return quorumMeta;
        }

        public String getQuorumName() {
            return quorumMeta.name;
        }

        public String getName() {
            return name;
        }
    }

    public static class QuorumMeta {
        private long id;
        private String name;

        public QuorumMeta(long quorumId) {
            this.id = quorumId;
            this.name = String.format("quorum-%d", quorumId);
        }

        public QuorumMeta(long quorumId, String name) {
            this.id = quorumId;
            this.name = name;
        }

        public long getId() {
            return id;
        }

        public String getName() {
            return name;
        }
    }

    class QuorumTrace {
        private String quorumIdStr;
        public int recorder;

        private TMB_Trace trace;
        private ReentrantReadWriteLock lock; // TODO: a


        QuorumTrace(long quorumId, TMB_Trace trace) { // make sure trace has events
            this.quorumIdStr = String.format("quorum-%d", quorumId);
            this.trace = trace;

            if (trace.getEvents().get(0).getService().equals(this.quorumIdStr)) {
                this.recorder = REQUESTER;
            } else {
                this.recorder = RESPONDER;
            }
        }
    }

    public class QuorumTraces {
        private QuorumMeta quorumMeta;
        private Map<Long, QuorumTrace> traces; // Map<traceId, Trace>
        private ReentrantReadWriteLock lock;

        QuorumTraces(QuorumMeta quorumMeta) {
            this.quorumMeta = quorumMeta;
            this.traces = new HashMap<>();
            this.lock = new ReentrantReadWriteLock();
        }

        public void printAllJSON() {
            lock.readLock().lock();
            for (Long key: traces.keySet()) {
                TMB_Helper.printf("[TMB_Store] [%s] %d: %d, %s\n", quorumMeta.getName(), key, traces.get(key).recorder, traces.get(key).trace.toJSON());
            }
            lock.readLock().unlock();
        }

        public void setTrace(TMB_Trace trace_) {
            lock.writeLock().lock();
            QuorumTrace quorumTrace = traces.get(trace_.getId());
            if (quorumTrace != null) {
                quorumTrace.trace.mergeEventsUnsafe(trace_.getEvents());
            } else {
                quorumTrace = new QuorumTrace(quorumMeta.getId(), trace_.copy());
            }
            traces.put(trace_.getId(), quorumTrace);
            lock.writeLock().unlock();
        }

        public void setTrace(TMB_Trace trace_, int newEvents) {
            lock.writeLock().lock();

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

    public void quorumQuit(QuorumMeta quorumMeta, TMB_Trace trace_) {
        quorumLock.readLock().lock();
        QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);
        quorumLock.readLock().unlock();

        QuorumTrace quorumTrace = quorumTraces.getQuorumTrace(trace_.getId());
        if (quorumTrace != null && quorumTrace.recorder == RESPONDER) {
            quorumTraces.removeQuorumTrace(trace_.getId());
        }
    }

    public void quorumSetTrace(ProcessorMeta procMeta, TMB_Trace trace_) {
        quorumSetTrace(procMeta.getQuorumMeta(), trace_);
    }

    public void quorumSetTrace(QuorumMeta quorumMeta, TMB_Trace trace_) {
        if (trace_.getId() == 0) {
            return;
        }
        QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);
        quorumTraces.setTrace(trace_);
    }

    public TMB_Trace quorumGetTrace(QuorumMeta quorumMeta, long traceId) {
        QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);

        quorumTraces.lock.readLock().lock();
        QuorumTrace quorumTrace = quorumTraces.getQuorumTrace(traceId);
        TMB_Trace trace = quorumTrace.trace;
        if (trace != null) {
            trace = trace.copy();
        }
        quorumTraces.lock.readLock().unlock();

        return trace;
    }

    public QuorumTraces getQuorumTraces(ProcessorMeta procMeta) {
        return getQuorumTraces(procMeta.getQuorumMeta());
    }

    public QuorumTraces getQuorumTraces(QuorumMeta quorumMeta) {
        long quorumId = quorumMeta.getId();

        quorumLock.readLock().lock();
        QuorumTraces quorumTraces = this.quorumTraces.get(quorumId);
        quorumLock.readLock().unlock();

        if (quorumTraces == null) {
            quorumLock.writeLock().lock();
            quorumTraces = this.quorumTraces.get(quorumId);
            if (quorumTraces == null) {
                quorumTraces = new QuorumTraces(quorumMeta);
                this.quorumTraces.put(quorumMeta.getId(), quorumTraces);
            }
            quorumLock.writeLock().unlock();
        }

        return quorumTraces;
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
            TMB_Helper.printf("[TMB_Store] unreachable code, trace: %s", trace.toJSON());
        }

        lock.writeLock().unlock();
    }

    private static Map<Long, TMB_Trace> thread_traces = new HashMap<>(); // TODO: a delete
    private static Map<Long, TMB_Trace> server_traces = new HashMap<>(); // TODO: a delete
    private static ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    public static void clearServerTraces() {
        lock.writeLock().lock();
        server_traces.clear();
        lock.writeLock().unlock();
    }

    public static void callerAppendEventsByThreadIdUnsafe(long threadId, TMB_Trace trace) {
        TMB_Trace trace_ = thread_traces.get(threadId);

        if (trace_ == null) {
            thread_traces.put(threadId, trace);
        } else {
            trace_.mergeEventsUnsafe(trace.getEvents());
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
}
