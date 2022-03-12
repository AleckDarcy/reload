package org.apache.zookeeper.trace;

import org.apache.jute.Record;

import java.util.ArrayList;

public class TMB_ClientPlugin {
    private TMB_Store.ProcessorMeta procMeta;
    private TMB_Trace trace;

    public TMB_ClientPlugin(int cnxnId) {
        int quorumId = - (Math.abs(cnxnId) + 1); // quorumId < 0
        this.procMeta = new TMB_Store.ProcessorMeta(new TMB_Store.QuorumMeta(quorumId, String.format("client%s", quorumId)), this);
    }

    public void TMB_Initialize(TMB_Trace trace_) {
        this.trace = trace_;
    }

    /**
     * Gets and deletes traces
     */
    public TMB_Trace TMB_Finalize() {
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
    public void outbound(Record request) {
        TMB_Trace trace = request.getTrace();
        // stub, should be called only once per client-level request
        // TODO: let client generate trace_id
        if (trace != null && trace.getId() > 0) {
            long id = TMB_Helper.newTraceId();
            trace.setId(id);
            trace.setEvents(new ArrayList<>(), 0);
            TMB_Helper.println("stub trace with id:" + id);

            trace.addEvent(this.procMeta, TMB_Event.Type.SERVICE_SEND, TMB_Helper.getClassNameFromObject(request), TMB_Helper.UUID());

            TMB_Store.getInstance().setTrace(this.procMeta, trace);
        }

        TMB_Helper.println("caller outbound ejects request: " + TMB_Helper.getClassNameFromObject(request) + "(" + TMB_Record.getString(request) + ")");
    }

    public void inbound(Record response) {
        TMB_Trace trace = response.getTrace();
        if (!trace.hasEvents()) {
            return;
        }

        String responseName = TMB_Helper.getClassNameFromObject(response);
        trace.addEvent(this.procMeta, TMB_Event.Type.SERVICE_RECV, responseName, trace.getEventUnsafe(0).getUuid());
        TMB_Store.getInstance().setTrace(procMeta, trace);

        TMB_Helper.println("caller inbound receives response: " + responseName + "(" + TMB_Record.getString(response) + ")");
    }
}
