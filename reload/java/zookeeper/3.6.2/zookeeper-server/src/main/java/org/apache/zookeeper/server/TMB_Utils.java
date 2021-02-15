package org.apache.zookeeper.server;

import org.apache.jute.Record;
import org.apache.zookeeper.server.quorum.QuorumPacket;
import org.apache.zookeeper.trace.TMB_Event;
import org.apache.zookeeper.trace.TMB_Helper;
import org.apache.zookeeper.trace.TMB_Store;
import org.apache.zookeeper.trace.TMB_Trace;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.Arrays;
import java.util.List;

public class TMB_Utils {
    public static final int EVENT_SERIALIZE_SIZE = 100;

    public static void printRequestForProcessor(String processorName, String quorumName, Object next, Request request) {
        String nextName = "null";
        if (next != null) {
            nextName = next.getClass().getCanonicalName();
        }

        Record record = request.record;
        Record txn = request.getTxn();

        String requestStr = String.format("(sessionid:0x%s, type:%s, cxid:0x%s, zxid:0x%s, txntype:%s, request:%s, record:%s, txn:%s)",
                Long.toHexString(request.sessionId),
                Request.op2String(request.type),
                Long.toHexString(request.cxid),
                Long.toHexString(request.getHdr() == null ? -2 : request.getHdr().getZxid()),
                request.getHdr() == null ? "unknown" : "" + request.getHdr().getType(),
                request.request == null ? "null": "valued",
                record == null ? "null": String.format("%s(trace:%s)", record.getClass().getCanonicalName(), record.getTrace() == null ? null: record.getTrace().getId() != 0),
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

    public static Record appendEvent(Record record, int type, String messageName, String service) {
        TMB_Trace trace = record.getTrace();
        if (trace == null) {
            return record;
        }

        List<TMB_Event> events = trace.getEvents();

        int eventSize = events.size();
        if (eventSize > 0) {
            TMB_Event lastEvent = events.get(eventSize - 1);
            String uuid = lastEvent.getUuid();
            if (type == TMB_Event.RECORD_FRWD) {
                uuid += "-0";
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

            trace.setEvents(events);
            record.setTrace(trace);
        }

        return record;
    }

    public static Record appendEvent(Record record, int type, String service) {
        return appendEvent(record, type, "", service);
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

            }
        }
    }
}
