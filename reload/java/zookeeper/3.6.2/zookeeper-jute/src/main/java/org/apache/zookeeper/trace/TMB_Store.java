package org.apache.zookeeper.trace;

import org.apache.jute.Record;
import org.apache.zookeeper.server.quorum.QuorumPacket;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class TMB_Store {
    private static TMB_Store instance = new TMB_Store();

    public static final int REQUESTER = 1;
    public static final int RESPONDER = 2;

    private TMB_Store() {
        this.quorums = new HashMap<>();
        this.lock = new ReentrantReadWriteLock();
    }

    public static TMB_Store getInstance() {
        return instance;
    }

    public static class ProcessorMeta {
        private QuorumMeta quorumMeta;
        private String name;

        public ProcessorMeta(QuorumMeta quorumMeta, Object processor) {
            this.quorumMeta = quorumMeta;
            this.name = TMB_Helper.getClassName(processor.getClass());
        }

        public QuorumMeta getQuorumMeta() {
            return this.quorumMeta;
        }

        public String getQuorumName() {
            return this.quorumMeta.name;
        }

        public String getName() {
            return this.name;
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
            return this.id;
        }

        public String getName() {
            return this.name;
        }
    }

    class QuorumTrace {
        private QuorumMeta quorumMeta;
        public int recorder;
        private TMB_Trace trace;

        QuorumTrace(QuorumMeta quorumMeta, TMB_Trace trace) { // make sure trace has events
            this.quorumMeta = quorumMeta;
            this.trace = trace;

            // TODO: a 1) get uuid of the last event; 2) find the first event of this uuid; 3) check this event
            if (trace.getEvents().get(0).getService().equals(quorumMeta.name)) {
                this.recorder = REQUESTER;
            } else {
                this.recorder = RESPONDER;
            }
        }

        public TMB_Trace getTraceCopyUnsafe() {
            if (this.trace == null) {
                return null;
            }

            return this.trace.copy();
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
            this.lock.readLock().lock();
            for (Long key: this.traces.keySet()) {
                TMB_Helper.printf("[TMB_Store] [%s] %d: %d, %s\n", this.quorumMeta.getName(), key, this.traces.get(key).recorder, this.traces.get(key).trace.toJSON());
            }
            this.lock.readLock().unlock();
        }

        public void setTrace(TMB_Trace trace_) {
            setTrace(trace_, trace_.getEvents().size());
        }

        // last newEvents events are new events
        public void setTrace(TMB_Trace trace_, int newEvents) {
            List<TMB_Event> events_ = trace_.getEvents();
            int eventSize = events_.size();
            newEvents = Math.min(eventSize, newEvents);
            if (newEvents == 0) {
                return;
            }
            List<TMB_Event> add = events_.subList(eventSize - newEvents, eventSize);

            this.lock.writeLock().lock();
            QuorumTrace quorumTrace = this.traces.get(trace_.getId());
            if (quorumTrace != null) {
                quorumTrace.trace.mergeEventsUnsafe(add);
            } else {
                quorumTrace = new QuorumTrace(this.quorumMeta, trace_.copy());
                this.traces.put(trace_.getId(), quorumTrace);
            }
            this.lock.writeLock().unlock();
        }

        public QuorumTrace getQuorumTrace(long traceId) {
            this.lock.readLock().lock();
            QuorumTrace quorumTrace = this.traces.get(traceId);
            this.lock.readLock().unlock();

            return quorumTrace;
        }

        public QuorumTrace getQuorumTraceUnsafe(long traceId) {
            return this.traces.get(traceId);
        }

        public QuorumTrace removeQuorumTrace(long traceId) {
            this.lock.readLock().lock();
            QuorumTrace quorumTrace = this.traces.remove(traceId);
            this.lock.readLock().unlock();

            return quorumTrace;
        }

        public QuorumTrace removeQuorumTraceUnsafe(long traceId) {
            return this.traces.remove(traceId);
        }
    }

    private Map<Long, QuorumTraces> quorums; // Map<quorumId, QuorumTrace>
    private ReentrantReadWriteLock lock;

    public void quit(QuorumMeta quorumMeta, TMB_Trace trace_) {
        this.lock.readLock().lock();
        QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);
        this.lock.readLock().unlock();

        QuorumTrace quorumTrace = quorumTraces.getQuorumTrace(trace_.getId());
        if (quorumTrace != null && quorumTrace.recorder == RESPONDER) {
            quorumTraces.removeQuorumTrace(trace_.getId());
        }
    }

    public void setTrace(ProcessorMeta procMeta, TMB_Trace trace_) {
        setTrace(procMeta.getQuorumMeta(), trace_);
    }

    public void setTrace(QuorumMeta quorumMeta, TMB_Trace trace_) {
        if (trace_.enabled()) {
            QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);
            quorumTraces.setTrace(trace_);
        }
    }

    public void setTrace(ProcessorMeta procMeta, TMB_Trace trace_, int newEvents) {
        setTrace(procMeta.getQuorumMeta(), trace_, newEvents);
    }

    public void setTrace(QuorumMeta quorumMeta, TMB_Trace trace_, int newEvents) {
        if (trace_.enabled()) {
            QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);
            quorumTraces.setTrace(trace_, newEvents);
        }
    }

    public TMB_Trace getTrace(ProcessorMeta procMeta, long traceId) {
        return getTrace(procMeta.getQuorumMeta(), traceId);
    }

    public TMB_Trace getTrace(QuorumMeta quorumMeta, long traceId) {
        QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);

        quorumTraces.lock.readLock().lock();
        QuorumTrace quorumTrace = quorumTraces.getQuorumTraceUnsafe(traceId);
        TMB_Trace trace = quorumTrace.getTraceCopyUnsafe();
        quorumTraces.lock.readLock().unlock();

        return trace;
    }

    public void removeTrace(ProcessorMeta procMeta, long traceId) {
        removeTrace(procMeta.getQuorumMeta(), traceId);
    }

    public void removeTrace(QuorumMeta quorumMeta, long traceId) {
        QuorumTraces quorumTraces = getQuorumTraces(quorumMeta);

        quorumTraces.lock.writeLock().lock();
        quorumTraces.removeQuorumTraceUnsafe(traceId);
        quorumTraces.lock.writeLock().unlock();
    }

    private QuorumTraces getQuorumTraces(QuorumMeta quorumMeta) {
        long quorumId = quorumMeta.getId();

        this.lock.readLock().lock();
        QuorumTraces quorum = this.quorums.get(quorumId);
        this.lock.readLock().unlock();

        if (quorum == null) {
            this.lock.writeLock().lock();
            quorum = this.quorums.get(quorumId);
            if (quorum == null) {
                quorum = new QuorumTraces(quorumMeta);
                this.quorums.put(quorumMeta.getId(), quorum);
            }
            this.lock.writeLock().unlock();
        }

        return quorum;
    }

    public void printQuorumTraces(QuorumMeta quorumMeta) {
        getQuorumTraces(quorumMeta).printAllJSON();
    }

    /**
     * Stores a trace associated with procMeta.
     * @param procMeta
     * @param record
     */
    public static void collectTrace(TMB_Store.ProcessorMeta procMeta, Record record) {
        TMB_Store.getInstance().setTrace(procMeta, record.getTrace());
    }

    /**
     * Called when quorum needs to capture events other than just TMB_EVENT.SERVICE_RECV.
     * @param procMeta
     * @param record
     * @param eventType
     */
    public static void collectTrace(TMB_Store.ProcessorMeta procMeta, Record record, int eventType) {
        TMB_Record.appendEvent(procMeta, record, eventType);
        collectTrace(procMeta, record);
    }

    // TODO: a all event are newly witnessed events
    /**
     * Called when quorum processes the incoming message.
     * Uses TMB_Event.SERVICE_RECV as default value of eventType.
     * @param procMeta
     * @param record    destination of deserialization
     * @param data      source of deserialization
     */
    public static void collectTrace(TMB_Store.ProcessorMeta procMeta, Record record, byte[] data) {
        if (data != null) {
            try {
                TMB_Record.deserialize(record, new ByteArrayInputStream(data));
                collectTrace(procMeta, record, TMB_Event.Type.SERVICE_RECV);
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    /**
     * Called when quorum processes the incoming message.
     * @param procMeta
     * @param record    destination of deserialization
     * @param qp        source of deserialization
     */
    public static void collectTrace(TMB_Store.ProcessorMeta procMeta, Record record, QuorumPacket qp) {
        collectTrace(procMeta, record, qp.getData());
    }
}
