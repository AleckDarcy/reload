package org.apache.zookeeper.server;

import org.apache.zookeeper.trace.TMB_Helper;

import java.util.Arrays;

public class TMB_Utils {
    public static void printRequestForProcessor(String processorName, String quorumName, Object next, Request request) {
        String nextName = "null";
        if (next != null) {
            nextName = next.getClass().getCanonicalName();
        }

        String requestStr = String.format("(sessionid:0x%s, type:%s, cxid:0x%s, zxid:0x%s, txntype:%s, request:%s)",
                Long.toHexString(request.sessionId),
                Request.op2String(request.type),
                Long.toHexString(request.cxid),
                Long.toHexString(request.getHdr() == null ? -2 : request.getHdr().getZxid()),
                request.getHdr() == null ? "unknown" : "" + request.getHdr().getType(),
                request.request == null ? "null": "valued");

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
}
