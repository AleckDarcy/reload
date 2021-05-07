package org.apache.zookeeper.server;

import org.apache.jute.Record;
import org.apache.zookeeper.proto.NullPointerResponse;
import org.apache.zookeeper.server.quorum.QuorumPacket;
import org.apache.zookeeper.trace.*;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Arrays;
import java.util.List;

public class TMB_Utils {
    public static final int EVENT_SERIALIZE_SIZE = 100;

    public static final String QUORUM_ACK = "QuorumAck";
    public static final String LEADER_COMMIT_READY = "LeaderCommitReady";
    public static final String LEADER_COMMIT = "LeaderCommit";
    public static final String LEADER_SYNC = "LeaderSync";

    public static void printRequestForProcessor(String processorName, TMB_Store.QuorumMeta quorumMeta, Object next, Request request) {
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

        TMB_Helper.printf(3, "[%s] %s, next %s, request-%d %s\n", quorumMeta.getName(), processorName, nextName, request.hashCode(), requestStr);
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

    public static NullPointerResponse commitHelperBegins(TMB_Store.QuorumMeta quorumMeta, Request request, String messageName, Class processor) {
        if (request != null) {
            Record txn = request.getTxn();
            if (txn != null) {
                TMB_Trace trace = txn.getTrace();
                if (trace != null && trace.getId() != 0) {
                    TMB_Trace trace_ = TMB_Store.getInstance().quorumGetTrace(quorumMeta, trace.getId());

                    if (trace_ != null) {
                        trace = trace_;
                    } else { // TODO: a report this case

                    }

                    // TODO: a uuid
                    List<TMB_Event> events = trace.getEvents();
                    int eventSize = events.size();
                    if (eventSize > 0) {
                        TMB_Event lastEvent = events.get(eventSize - 1);
                        String uuid = TMB_Helper.UUID();

                        TMB_Event event = new TMB_Event(TMB_Event.LOGIC_COMMIT_READY, TMB_Helper.currentTimeNanos(), LEADER_COMMIT_READY, uuid, quorumMeta.getName(), processor);
                        trace.addEvent(event);
                    }

                    return new NullPointerResponse(messageName, trace);
                }
            }
        }

        return null;
    }

    public static void commitHelperEnds(TMB_Store.QuorumMeta quorumMeta, Record record) {
        if (record != null) {
            TMB_Trace trace = record.getTrace();
            if (trace != null && trace.getId() != 0) {
                TMB_Store.getInstance().quorumSetTrace(quorumMeta, trace);
            }
        }
    }

    public static byte[] ackHelper(Request request, String messageName, TMB_Store.QuorumMeta quorumMeta, Class processor) throws FaultInjectedException {
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
                    record = TMB_Utils.appendEvent(record, TMB_Event.SERVICE_SEND, messageName, quorumMeta, false, processor);

                    try {
                        data = TMB_Helper.serialize(record);
                    } catch (IOException e) {
                    }
                }
            }
        }

        return data;
    }

    public static Record pRequestHelper(Record record, Record txn) {
        // TODO: a capture RECV event
        txn.setTrace(record.getTrace());

        return txn;
    }

    public static Record appendEvent(Record record, int type, String messageName, TMB_Store.QuorumMeta quorumMeta, boolean truncateUUID, Class processor) {
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
            events.add(new TMB_Event(type, TMB_Helper.currentTimeNanos(), messageName, uuid, quorumMeta.getName(), processor));

            trace.setEvents(events, 1);
            record.setTrace(trace);
        }

        return record;
    }

    public static Record appendEvent(Record record, int type, TMB_Store.QuorumMeta quorumMeta, Class processor) {
        return appendEvent(record, type, "", quorumMeta, false, processor);
    }

    // before sending the message,
    // deserialize message and append an event to the message,
    // finally serialize message
    public static ByteBuffer appendEvent(ByteBuffer data, Record request, int type, TMB_Store.QuorumMeta quorumMeta, Class processor) throws IOException {
        try {
            ByteBufferInputStream.byteBuffer2Record(data, request);
        } catch (IOException e) {
            data.rewind();
            throw e;
        }


        request = appendEvent(request, type, quorumMeta, processor);

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

    // TODO: a using int[] types
    public static ByteBuffer appendEvents(ByteBuffer data, Record request, int type1, int type2, TMB_Store.QuorumMeta quorumMeta, Class processor) throws IOException {
        try {
            ByteBufferInputStream.byteBuffer2Record(data, request);
        } catch (IOException e) {
            data.rewind();
            throw e;
        }


        request = appendEvent(request, type1, quorumMeta, processor);
        request = appendEvent(request, type2, quorumMeta, processor);

        TMB_Trace trace = request.getTrace();
        if (trace != null) {
            List<TMB_Event> events = request.getTrace().getEvents();

            int eventSize = events.size();
            if (eventSize > 0) {
                ByteBuffer bb = ByteBuffer.allocate(data.capacity() + EVENT_SERIALIZE_SIZE * 2);
                ByteBufferOutputStream.record2ByteBuffer(request, bb);

                return bb;
            }
        }

        data.rewind();

        return data;
    }

    public static void quorumCollectTraceFromQuorumPacket(TMB_Store.QuorumMeta quorumMeta, Record record, QuorumPacket qp, Class processor) {
        byte[] data = qp.getData();
        if (data != null) {
            try {
                TMB_Helper.deserialize(new ByteArrayInputStream(data), record);
                record = TMB_Utils.appendEvent(record, TMB_Event.SERVICE_RECV, quorumMeta, processor);
                TMB_Store.getInstance().quorumSetTrace(quorumMeta, record.getTrace());
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
