package org.apache.zookeeper.server;

import org.apache.jute.BinaryOutputArchive;
import org.apache.jute.Record;
import org.apache.zookeeper.proto.NullPointerResponse;
import org.apache.zookeeper.server.quorum.QuorumPacket;
import org.apache.zookeeper.trace.*;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.List;

public class TMB_Utils {
    public static final int EVENT_SERIALIZE_SIZE = 100;

    public static final String QUORUM_ACK = "QuorumAck";
    public static final String LEADER_PRPS_READY = "LeaderProposeReady";
    public static final String LEADER_ENOUGH_ACK = "LeaderEnoughACK";
    public static final String LEADER_COMMIT_READY = "LeaderCommitReady";
    public static final String LEADER_COMMIT = "LeaderCommit";
    public static final String LEADER_SYNC = "LeaderSync";

    // an extension of org.apache.zookeeper.server.Request
    // Request.request will be deserialized once a life time inside a particular quorum
    // all field un-nullable
    public static class RequestExt {
        private boolean traced;
        private ProcessorFlag procFlag;
        private Record message;
        private String uuid;

        public RequestExt(Record message, ProcessorFlag procFlag) {
            TMB_Trace trace = message.getTrace();
            if (trace != null && trace.getId() != 0) {
                this.traced = true;
                this.message = message;
            } else {
                this.traced = false;
            }

            this.procFlag = new ProcessorFlag(procFlag);
        }

        public boolean isTraced() {
            return this.traced;
        }

        public ProcessorFlag getProcessorFlag() {
            return this.procFlag;
        }

        public void updateProcessorFlag(ProcessorFlag procFlag) {
            if (this.procFlag.isNull()) {
                this.procFlag = new ProcessorFlag(procFlag);
            } else {
                this.procFlag.update(procFlag);
            }
        }

        public Record getMessage() {
            return this.message;
        }

        public void setUUID(String uuid) {
            this.uuid = uuid;
        }

        public String getUUID() {
            return this.uuid;
        }
    }

    public static class ProcessorFlag {
        public static final ProcessorFlag NULL = new ProcessorFlag(0x0);
        public static final ProcessorFlag RECV = new ProcessorFlag(0x1);

        private int flag;

        private ProcessorFlag(int flag) { this.flag = flag; }

        public ProcessorFlag() { this.flag = NULL.flag; }

        public ProcessorFlag(ProcessorFlag flag) { this.flag = flag.flag; }

        public void update(ProcessorFlag flag) { this.flag |= flag.flag; }

        public boolean isNull() { return flag == NULL.flag; }

        public boolean isReceived() { return (flag & RECV.flag) != 0; }
    }

    public static void processorPrintsRequest(TMB_Store.ProcessorMeta procMeta, String info, Object next, Request request) {
        if (!TMB_Helper.printable) {
            return;
        }

        String nextName = "null";
        if (next != null) {
            nextName = TMB_Helper.getClassNameFromObject(next);
        }

        Record txn = request.getTxn();

        String requestStr = String.format("(sessionid:0x%s,type:%s,cxid:0x%s,zxid:0x%s,txntype:%s,request:%s,record:%s)",
                Long.toHexString(request.sessionId),
                Request.op2String(request.type),
                Long.toHexString(request.cxid),
                Long.toHexString(request.getHdr() == null ? -2 : request.getHdr().getZxid()),
                request.getHdr() == null ? "unknown" : "" + request.getHdr().getType(),
                request.request == null ? "null" : "valued",
                txn == null ? "null" : String.format("%s-%d(trace:%s)", TMB_Helper.getClassNameFromObject(txn), txn.hashCode(), txn.getTrace() == null ? null : txn.getTrace().getId() != 0));

        if (info != null && info.length() != 0) {
            TMB_Helper.printf(procMeta, 3, "%s, next:%s, request-%d:%s\n", info, nextName, request.hashCode(), requestStr);
        } else {
            TMB_Helper.printf(procMeta, 3, "next:%s, request-%d:%s\n", nextName, request.hashCode(), requestStr);
        }
    }

    public static NullPointerResponse processAckHelperBegins(TMB_Store.ProcessorMeta procMeta, byte[] data) {
        NullPointerResponse record = new NullPointerResponse();
        quorumCollectTrace(procMeta, record, data);

        return record;
    }

    // leader processes request as an ACK
    public static void processAckHelperBegins(TMB_Store.ProcessorMeta procMeta, Record record) {
        if (record == null) {
            return;
        }

        TMB_Trace trace = record.getTrace();
        if (trace == null || trace.getId() == 0) {
            return;
        }

        TMB_Event event = new TMB_Event();
        quorumCollectTrace(procMeta, record, TMB_Event.SERVICE_RECV);
    }

    public static NullPointerResponse commitHelperBegins(TMB_Store.ProcessorMeta procMeta, Request request, String messageName) {
        if (request != null) {
            Record txn = request.getTxn();
            if (txn != null) {
                TMB_Trace trace = txn.getTrace();
                if (trace != null && trace.getId() != 0) {
                    TMB_Trace trace_ = TMB_Store.getInstance().quorumGetTrace(procMeta.getQuorumMeta(), trace.getId());

                    if (trace_ != null) {
                        trace = trace_;
                    } else { // TODO: a report this case

                    }

                    return new NullPointerResponse(messageName, trace);
                }
            }
        }

        return null;
    }

    public static void commitHelperEnds(TMB_Store.ProcessorMeta procMeta, Record record) {
        if (record != null) {
            TMB_Trace trace = record.getTrace();
            if (trace != null && trace.getId() != 0) {
                TMB_Store.getInstance().quorumSetTrace(procMeta, trace);
            }
        }
    }

    public static byte[] ackHelper(TMB_Store.ProcessorMeta procMeta, Request request, String messageName) throws FaultInjectedException {
        byte[] data = null;
        if (request != null) {
            Record txn = request.getTxn();
            if (txn != null) {
                Record record = new NullPointerResponse(messageName);
                TMB_Trace trace = txn.getTrace();
                if (trace != null && trace.getId() != 0) {
                    // Direct Response Circle, the trace is carried by txn
                    // do not need TMB_Store.getInstance().quorumGetTrace(quorumId, trace.getId()) here

                    TMB_Helper.checkTFIs(trace, messageName);

                    record.setTrace(trace);
                    record = TMB_Utils.appendEvent(procMeta, record, TMB_Event.SERVICE_SEND, messageName);

                    try {
                        data = TMB_Helper.serialize(record);
                    } catch (IOException e) {
                    }
                }
            }
        }

        return data;
    }

    // TODO: a delete
    public static Record pRequestHelper(TMB_Store.ProcessorMeta procMeta, Record record, Record txn) {
        TMB_Trace trace = record.getTrace();
        if (trace != null && trace.getId() != 0) {
            List<TMB_Event> events = trace.getEvents();
            if (events.size() != 0) {
                TMB_Event lastEvent = events.get(events.size() - 1);
                TMB_Event event = new TMB_Event(TMB_Event.SERVICE_RECV, lastEvent.getMessage_name(), lastEvent.getUuid(), procMeta);
                events.add(event);
                trace.setEvents(events);
            }
        }
        txn.setTrace(record.getTrace());

        return txn;
    }

    public static void pRequestHelper(TMB_Store.ProcessorMeta procMeta, Request request, Record record, Record txn) {
        TMB_Trace trace = record.getTrace();
        if (trace != null && trace.getId() != 0) {
            List<TMB_Event> events = trace.getEvents();
            if (events.size() != 0) {
                TMB_Event lastEvent = events.get(events.size() - 1);
                TMB_Event event = new TMB_Event(TMB_Event.SERVICE_RECV, lastEvent.getMessage_name(), lastEvent.getUuid(), procMeta);
                events.add(event);
                trace.setEvents(events);
            }
            RequestExt requestExt = new RequestExt(record, ProcessorFlag.RECV);
            request.setRequestExt(requestExt);
        }
        txn.setTrace(trace);
        request.setTxn(txn);
    }

    public static Record appendEvent(TMB_Store.ProcessorMeta procMeta, Record record, int type, String messageName) {
        TMB_Trace trace = record.getTrace();
        if (trace == null && trace.getId() == 0) {
            return record;
        }

        List<TMB_Event> events = trace.getEvents();

        int eventSize = events.size();
        if (eventSize > 0) {
            TMB_Event lastEvent = events.get(eventSize - 1);
            String uuid = lastEvent.getUuid();
            if (messageName.equals("")) {
                messageName = lastEvent.getMessage_name();
            }
            events.add(new TMB_Event(type, messageName, uuid, procMeta));

            trace.setEvents(events, 1);
            record.setTrace(trace);
        }

        return record;
    }

    public static Record appendEvent(TMB_Store.ProcessorMeta procMeta, Record record, int type) {
        return appendEvent(procMeta, record, type, "");
    }

    // TODO: a check request.getTxn() is always null
    public static void forwardHelper(TMB_Store.ProcessorMeta procMeta, Request request, Record record) throws IOException {
        ByteBuffer data = request.request;
        try {
            ByteBufferInputStream.byteBuffer2Record(data, record);
        } catch (IOException e) {
            data.rewind();
            throw e;
        }

        TMB_Trace trace = record.getTrace();
        if (trace != null) {
            request.setRequestExt(new RequestExt(record, ProcessorFlag.RECV));
            List<TMB_Event> events = trace.getEvents();
            int eventSize = events.size();
            if (eventSize > 0) {
                TMB_Event lastEvent = events.get(eventSize - 1);

                events.add(new TMB_Event(TMB_Event.SERVICE_RECV, lastEvent.getMessage_name(), lastEvent.getUuid(), procMeta));
                events.add(new TMB_Event(TMB_Event.SERVICE_FRWD, lastEvent.getMessage_name(), TMB_Helper.UUID(), procMeta));

                ByteArrayOutputStream baos = new ByteArrayOutputStream(data.capacity() + EVENT_SERIALIZE_SIZE * 2);
                BinaryOutputArchive boa = BinaryOutputArchive.getArchive(baos);
                record.serialize(boa, "request");

                request.request = ByteBuffer.wrap(baos.toByteArray());

                return;
            }
        }

        data.rewind();
    }

    // default event: SERVICE_RECV
    public static void quorumCollectTrace(TMB_Store.ProcessorMeta procMeta, Record record) {
        TMB_Store.getInstance().quorumSetTrace(procMeta, record.getTrace());
    }

    public static void quorumCollectTrace(TMB_Store.ProcessorMeta procMeta, Record record, int eventType) {
        record = TMB_Utils.appendEvent(procMeta, record, eventType);
        quorumCollectTrace(procMeta, record);
    }

    public static void quorumCollectTrace(TMB_Store.ProcessorMeta procMeta, Record record, byte[] data) {
        if (data != null) {
            try {
                TMB_Helper.deserialize(new ByteArrayInputStream(data), record);
                quorumCollectTrace(procMeta, record, TMB_Event.SERVICE_RECV);
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

    public static void quorumCollectTrace(TMB_Store.ProcessorMeta procMeta, Record record, QuorumPacket qp) {
        quorumCollectTrace(procMeta, record, qp.getData());
    }
}
