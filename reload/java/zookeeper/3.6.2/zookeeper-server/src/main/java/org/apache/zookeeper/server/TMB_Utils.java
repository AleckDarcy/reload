package org.apache.zookeeper.server;

import org.apache.jute.BinaryOutputArchive;
import org.apache.jute.Record;
import org.apache.zookeeper.proto.NullPointerResponse;
import org.apache.zookeeper.trace.*;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.List;

public class TMB_Utils {
    public static final int EVENT_SERIALIZE_SIZE = 100;

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
            if (trace.hasEvents()) {
                this.traced = true;
                this.message = message;

                if (procFlag == ProcessorFlag.RECV) {
                    this.uuid = trace.getLastEvent().getUuid();
                }
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

        private final TMB_Trace disabledTrace = new TMB_Trace();

        public TMB_Trace getTrace() {
            if (this.message == null) {
                return disabledTrace;
            }

            return this.message.getTrace();
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

        public boolean isNull() { return this.flag == NULL.flag; }

        public boolean isReceived() { return (this.flag & RECV.flag) != 0; }
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
                txn == null ? "null" : String.format("%s-%d(trace:%s)", TMB_Helper.getClassNameFromObject(txn), txn.hashCode(), txn.getTrace().hasEvents()));

        if (info != null && info.length() != 0) {
            TMB_Helper.printf(procMeta, 3, "%s, next:%s, request-%d:%s\n", info, nextName, request.hashCode(), requestStr);
        } else {
            TMB_Helper.printf(procMeta, 3, "next:%s, request-%d:%s\n", nextName, request.hashCode(), requestStr);
        }
    }

    /**
     * Captures events (SERVICE_RECV or PROCESSOR_RECV) to and stores records to RequestExt and Store
     * @param procMeta
     * @param request
     */
    public static void processRequestHelperBegins(TMB_Store.ProcessorMeta procMeta, Request request) {
        TMB_Utils.RequestExt requestExt = request.getRequestExt();
        if (requestExt == null || !requestExt.traced) {
            return;
        }

        TMB_Trace trace = requestExt.getTrace();
        if (!trace.hasEvents()) {
            return;
        }

        Record txn = request.getTxn();
        if (txn == null) {
            return;
        }

        int eventType = TMB_Event.Type.SERVICE_RECV;
        TMB_Utils.ProcessorFlag procFlag = request.getProcessorFlag();
        if (procFlag.isReceived()) {
            eventType = TMB_Event.Type.PROCESSOR_RECV;
        } else {
            requestExt.updateProcessorFlag(TMB_Utils.ProcessorFlag.RECV);
        }

        String requestName = TMB_Helper.getClassNameFromObject(txn);
        TMB_Event preEvent = trace.getEventUnsafe(0); // TODO: （a）get uuid from RequestExt
        trace.addEvent(procMeta, eventType, requestName, preEvent.getUuid());

        TMB_Store.getInstance().setTrace(procMeta, trace);
    }

    /**
     * Assigns events captured so far (from Request and a new SERVICE_SEND event) to the response
     * @param procMeta
     * @param request
     * @param response
     */
    public static void processRequestHelperEnds(TMB_Store.ProcessorMeta procMeta, Request request, Record response) {
        RequestExt requestExt = request.getRequestExt();
        if (requestExt == null) {
            return;
        }
        TMB_Trace trace = requestExt.getTrace();
        if (!trace.hasEvents()) {
            return;
        }

        TMB_Trace trace_ = TMB_Store.getInstance().getTrace(procMeta, trace.getId());
        trace_.addEvent(procMeta, TMB_Event.Type.SERVICE_SEND, TMB_Helper.getClassNameFromObject(response), requestExt.getUUID());
        response.setTrace(trace_);
        // TODO: (a) delete trace from TMB_Store
    }

    public static NullPointerResponse processAckHelperBegins(TMB_Store.ProcessorMeta procMeta, byte[] data) {
        NullPointerResponse record = new NullPointerResponse();
        TMB_Store.collectTrace(procMeta, record, data);

        return record;
    }

    public static NullPointerResponse commitHelperBegins(TMB_Store.ProcessorMeta procMeta, Request request, String messageName) {
        if (request != null) {
            Record txn = request.getTxn();
            if (txn != null) {
                TMB_Trace trace = txn.getTrace();
                if (trace.enabled()) {
                    TMB_Trace trace_ = TMB_Store.getInstance().getTrace(procMeta.getQuorumMeta(), trace.getId());

                    if (trace_ != null) {
                        trace = trace_;
                    } else { // TODO: (a) report this case

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
            if (trace.hasEvents()) {
                TMB_Store.getInstance().setTrace(procMeta, trace);
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
                if (trace.hasEvents()) {
                    // DRC, the trace is carried by txn
                    trace.checkTFIs(messageName);
                    trace.addEvent(procMeta, TMB_Event.Type.SERVICE_SEND, messageName);
                    record.setTrace(trace);

                    // TODO: (a) check txn.getTrace() == record.getTrace()
                    try {
                        data = TMB_Record.serialize(record);
                    } catch (IOException e) {
                    }
                }
            }
        }

        return data;
    }

    public static void pRequestHelper(TMB_Store.ProcessorMeta procMeta, Request request, Record record, Record txn) {
        TMB_Trace trace = record.getTrace();
        if (trace.hasEvents()) {
            trace.addEvent(procMeta, TMB_Event.Type.SERVICE_RECV);
            request.setRequestExt(new RequestExt(record, ProcessorFlag.RECV));
            txn.setTrace(trace);
        }
        request.setTxn(txn);
    }

    // TODO: (a) check request.getTxn() is always null
    public static void forwardHelper(TMB_Store.ProcessorMeta procMeta, Request request, Record record) throws IOException {
        ByteBuffer data = request.request;
        try {
            ByteBufferInputStream.byteBuffer2Record(data, record);
        } catch (IOException e) {
            throw e;
        } finally {
            data.rewind();
        }

        TMB_Trace trace = record.getTrace();
        if (trace.hasEvents()) {
            request.setRequestExt(new RequestExt(record, ProcessorFlag.RECV)); // TODO: (a) initialize RequestExt elsewhere

            TMB_Event lastEvent = trace.getLastEvent();
            // TODO: (a) find a better place (if could) to capture SERVICE_RECV.
            trace.addEvent(procMeta, TMB_Event.Type.SERVICE_RECV, lastEvent.getMessage_name(), lastEvent.getUuid());
            trace.addEvent(procMeta, TMB_Event.Type.SERVICE_FRWD, lastEvent.getMessage_name(), TMB_Helper.UUID());
            TMB_Store.getInstance().setTrace(procMeta, trace, 2);

            ByteArrayOutputStream baos = new ByteArrayOutputStream(data.capacity() + EVENT_SERIALIZE_SIZE * 2);
            BinaryOutputArchive boa = BinaryOutputArchive.getArchive(baos);
            record.serialize(boa, "request");

            request.request = ByteBuffer.wrap(baos.toByteArray());
        }
    }
}
