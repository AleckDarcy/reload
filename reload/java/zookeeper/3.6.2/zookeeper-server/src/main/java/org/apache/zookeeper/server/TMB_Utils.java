package org.apache.zookeeper.server;

import org.apache.jute.Record;
import org.apache.zookeeper.proto.NullPointerResponse;
import org.apache.zookeeper.server.quorum.QuorumPacket;
import org.apache.zookeeper.trace.*;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Arrays;
import java.util.List;

public class TMB_Utils {
    public static final int EVENT_SERIALIZE_SIZE = 100;

    public static final String QUORUM_ACK = "QuorumAck";
    public static final String LEADER_COMMIT = "LeaderCommit";
    public static final String LEADER_SYNC = "LeaderSync";

    public static void printRequestForProcessor(String processorName, String quorumName, Object next, Request request) {
        String nextName = "null";
        if (next != null) {
            nextName = next.getClass().getCanonicalName();
        }

        Record txn = request.getTxn();

        String requestStr = String.format("(sessionid:0x%s, type:%s, cxid:0x%s, zxid:0x%s, txntype:%s, request:%s, record:%s)",
                Long.toHexString(request.sessionId),
                Request.op2String(request.type),
                Long.toHexString(request.cxid),
                Long.toHexString(request.getHdr() == null ? -2 : request.getHdr().getZxid()),
                request.getHdr() == null ? "unknown" : "" + request.getHdr().getType(),
                request.request == null ? "null": "valued",
                txn == null ? "null": String.format("%s(trace:%s)", txn.getClass().getCanonicalName(), txn.getTrace() == null ? null: txn.getTrace().getId() != 0));

        TMB_Helper.printf(3, "[%s] %s, next %s, request-%d %s\n", quorumName, processorName, nextName, request.hashCode(), requestStr);
    }

    public static void printRequestForProcessorUnsafe(String processorName, String quorumName, Object next, Request request) {
        String nextName = "null";
        if (next != null) {
            nextName = next.getClass().getCanonicalName();
        }
        String requestStr = "null";
        if (request.request != null) {
            requestStr = Arrays.toString(request.request.array());
            request.request.rewind();
        }

        TMB_Helper.printf(3, "[%s] %s, next %s, request-%d %s\n", quorumName, processorName, nextName, request.hashCode(), requestStr);
    }

    public static byte[] commitHelper(Request request, String messageName, String quorumName, int quorumId) throws FaultInjectedException {
        byte[] data = null;
        if (request != null) {
            Record txn = request.getTxn();
            if (txn != null) {
                Record record = new NullPointerResponse(messageName);
                TMB_Trace trace = txn.getTrace();
                if (trace != null && trace.getId() != 0) {
                    TMB_Trace trace_ = TMB_Store.getInstance().quorumGetTrace(quorumId, trace.getId());
                    TMB_Helper.checkTFIs(trace_, messageName);

                    record.setTrace(trace_);
                    record = TMB_Utils.appendEvent(record, TMB_Event.RECORD_SEND, messageName, quorumName, true);

                    try {
                        ByteArrayOutputStream bao = TMB_Helper.serialize(record);
                        data = bao.toByteArray();
                    } catch (IOException e) {
                    }
                }
            }
        }

        return data;
    }

    public static byte[] ackHelper(Request request, String messageName, String quorumName, int quorumId) throws FaultInjectedException {
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
                    record = TMB_Utils.appendEvent(record, TMB_Event.RECORD_SEND, messageName, quorumName, false);

                    try {
                        ByteArrayOutputStream bao = TMB_Helper.serialize(record);
                        data = bao.toByteArray();
                    } catch (IOException e) {
                    }
                }
            }
        }

        return data;
    }

    public static Record pRequestHelper(Record record, Record txn) {
        txn.setTrace(record.getTrace());

        return txn;
    }

    public static Record appendEvent(Record record, int type, String messageName, String service, boolean truncateUUID) {
        TMB_Trace trace = record.getTrace();
        if (trace == null) {
            return record;
        }

        List<TMB_Event> events = trace.getEvents();

        int eventSize = events.size();
        if (eventSize > 0) {
            TMB_Event lastEvent = events.get(eventSize - 1);
            String uuid = lastEvent.getUuid();
            if (truncateUUID && uuid.length() == TMB_Helper.UUID_LEN + TMB_Helper.UUID_SUF_LEN) {
                uuid = uuid.substring(0, TMB_Helper.UUID_LEN);
            }

            if (messageName.equals("")) {
                messageName = lastEvent.getMessage_name();
            }
            events.add(new TMB_Event(
                    type,
                    TMB_Helper.currentTimeNanos(),
                    messageName,
                    uuid,
                    service));

            trace.setEvents(events, 1);
            record.setTrace(trace);
        }

        return record;
    }

    public static Record appendEvent(Record record, int type, String service) {
        return appendEvent(record, type, "", service, false);
    }

    // before sending the message,
    // deserialize message and append an event to the message,
    // finally serialize message
    public static ByteBuffer appendEvent(ByteBuffer data, Record request, int type, String service) throws IOException {
        try {
            ByteBufferInputStream.byteBuffer2Record(data, request);
        } catch (IOException e) {
            data.rewind();
            throw e;
        }


        request = appendEvent(request, type, service);

        TMB_Trace trace = request.getTrace();
        if (trace != null) {
            List<TMB_Event> events = request.getTrace().getEvents();

            int eventSize = events.size();
            if (eventSize > 0) {
                ByteBuffer bb = ByteBuffer.allocate(data.capacity() + EVENT_SERIALIZE_SIZE);
                ByteBufferOutputStream.record2ByteBuffer(request, bb);

                return bb;
            }
        }

        data.rewind();

        return data;
    }

    public static void quorumCollectTraceFromQuorumPacket(long quorumId, String quorumName, Record record, QuorumPacket qp) {
        byte[] data = qp.getData();
        if (data != null) {
            try {
                TMB_Helper.deserialize(new ByteArrayInputStream(data), record);
                record = TMB_Utils.appendEvent(record, TMB_Event.RECORD_RECV, quorumName);
                TMB_Store.getInstance().quorumSetTrace(quorumId, record.getTrace());
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
