package org.apache.zookeeper.trace;

import org.apache.jute.Record;

import java.util.ArrayList;
import java.util.List;

public class TMB_ClientPlugin {
    private TMB_Store.ProcessorMeta procMeta;
    private TMB_Trace trace;

    public TMB_ClientPlugin(int cnxnId) {
        int quorumId = - (Math.abs(cnxnId) + 1); // quorumId < 0
        this.procMeta = new TMB_Store.ProcessorMeta(new TMB_Store.QuorumMeta(quorumId, String.format("client%s", quorumId)), this);
    }

    public void TMBInitialize(TMB_Trace trace_) {
        this.trace = trace_;
    }

    /**
     * Gets and deletes traces
     */
    public TMB_Trace TMBFinalize() {
        TMB_Trace trace = TMB_Store.getInstance().getTrace(this.procMeta, this.trace.getId());
        if (trace != null) {
            TMB_Helper.printf(this.procMeta, "trace: %s\n", trace.toJSON());
            TMB_Store.getInstance().removeTrace(this.procMeta, trace.getId());
        }

        return trace;
    }

    public TMB_Trace getTrace() {
        return this.trace;
    }

    /**
     * submitRequest -> generate request -> callerOutbound -> network -> callerInbound -> process response
     */
    public void callerOutbound(String service, Record request) {
        TMB_Trace trace = request.getTrace();
        // stub, should be called only once per client-level request
        // TODO: let client generate trace_id
        if (trace.enabled()) {
            long id = TMB_Helper.newTraceId();
            trace.setId(id);
            trace.setEvents(new ArrayList<>(), 0);
            TMB_Helper.println("stub trace with id:" + id);
        }

        if (trace.enabled()) {
            String requestName = TMB_Helper.getClassNameFromObject(request);
            String uuid = TMB_Helper.UUID();
            TMB_Event event = new TMB_Event(TMB_Event.Type.SERVICE_SEND, requestName, uuid, service, TMB_ClientPlugin.class);

            List<TMB_Event> events = trace.getEvents();
            events.add(event);
            trace.setEvents(events, 1);

            TMB_Store.getInstance().setTrace(this.procMeta, trace);
        }

        TMB_Helper.println("caller outbound ejects request: " + TMB_Helper.getClassNameFromObject(request) + "(" + TMB_Record.getString(request) + ")");
    }

    public void callerInbound(String service, Record response) {
        TMB_Trace trace = response.getTrace();
        if (!trace.enabled()) {
            return;
        }

        String responseName = TMB_Helper.getClassNameFromObject(response);
        List<TMB_Event> events = trace.getEvents();
        if (events.size() == 0) {
            TMB_Helper.println("caller inbound receives trace without events: " + responseName + "(" + TMB_Record.getString(response) + ")");
            return;
        }

        String uuid = events.get(0).getUuid();
        TMB_Event event = new TMB_Event(TMB_Event.Type.SERVICE_RECV, responseName, uuid, service, TMB_ClientPlugin.class);
        trace.addEvent(event);

        TMB_Store.getInstance().setTrace(procMeta, trace);

        TMB_Helper.println("caller inbound receives response: " + responseName + "(" + TMB_Record.getString(response) + ")");
    }
}
