package org.apache.zookeeper.trace;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class TMB_ThreadTraces {
    private Map<UUID, TMB_Trace> traces;

    public TMB_ThreadTraces() {
        this.traces = new HashMap<UUID, TMB_Trace>();
    }

    public void put(UUID uuid, TMB_Trace trace) {
        this.traces.put(uuid, trace);
    }

    public TMB_Trace get(UUID uuid) {
        return this.traces.get(uuid);
    }

    public void remove(UUID uuid) {
        this.traces.remove(uuid);
    }
}