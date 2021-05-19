package org.apache.zookeeper.trace;

import org.apache.jute.Record;

import java.util.ArrayList;
import java.util.List;

public class TMB_ClientPlugin {
    private final long threadID = Thread.currentThread().getId(); // TODO: a delete

    private TMB_Store.ProcessorMeta procMeta;
    private TMB_Trace trace;

    /**
     *
     */
    public void TMBInitialize(TMB_Trace trace_) {
        trace = trace_;
        long quorumID = -Thread.currentThread().getId();
        procMeta = new TMB_Store.ProcessorMeta(new TMB_Store.QuorumMeta(quorumID, String.format("client%s", quorumID)), this);
    }

    /**
     * Get and delete traces
     */
    public TMB_Trace TMBFinalize() {
        TMB_Trace trace = TMB_Store.getByThreadId(threadID);

        if (trace != null) {
            TMB_Helper.println("Trace for thread " + threadID + ": " + trace.toJSON());
            TMB_Store.removeByThreadId(threadID);
        }

        return trace;
    }

    public TMB_Trace getTrace() {
        return trace;
    }

    /**
     * submitRequest -> generate request -> callerOutbound -> network -> callerInbound -> process response
     */
    public static void callerOutbound(String service, Record request) {
        TMB_Trace trace = request.getTrace();
        // stub, should be called only once per client-level request
        // TODO: let client generate trace_id
        if (trace.getId() == 0) {
            long id = TMB_Helper.newTraceId();
            trace.setId(id);
            trace.setEvents(new ArrayList<>(), 0);
            TMB_Helper.println("stub trace with id:" + id);
        }

        if (trace.getId() != 0) {
            long threadId = Thread.currentThread().getId();
            String requestName = TMB_Helper.getClassNameFromObject(request);
            String uuid = TMB_Helper.UUID();
            TMB_Event event = new TMB_Event(TMB_Event.SERVICE_SEND, requestName, uuid, service, TMB_ClientPlugin.class);

            List<TMB_Event> events = trace.getEvents();
            events.add(event);
            trace.setEvents(events, 1);

            TMB_Store.getInstance().callerAppendEventsByThreadIdUnsafe(threadId, trace);
        }

        TMB_Helper.println("caller outbound ejects request: " + TMB_Helper.getClassNameFromObject(request) + "(" + TMB_Helper.getString(request) + ")");
    }

    public static void callerInbound(String service, Record response) {
        TMB_Trace trace = response.getTrace();
        if (trace.getId() == 0) {
            return;
        }

        String responseName = TMB_Helper.getClassNameFromObject(response);
        List<TMB_Event> events = trace.getEvents();
        if (events.size() == 0) {
            TMB_Helper.println("caller inbound receives trace without events: " + responseName + "(" + TMB_Helper.getString(response) + ")");

            return;
        }

        long threadId = Thread.currentThread().getId();
        String uuid = events.get(0).getUuid();
        TMB_Event event = new TMB_Event(TMB_Event.SERVICE_RECV, responseName, uuid, service, TMB_ClientPlugin.class);
        trace.addEvent(event);

        TMB_Store.getInstance().callerAppendEventsByThreadIdUnsafe(threadId, trace);

        TMB_Helper.println("caller inbound receives response: " + responseName + "(" + TMB_Helper.getString(response) + ")");
    }
}
