package org.apache.zookeeper.trace;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class _3MB_ThreadTraces {
    private Map<UUID, _3MB_Trace> traces;

    public _3MB_ThreadTraces() {
        this.traces = new HashMap<UUID, _3MB_Trace>();
    }

    public void put(UUID uuid, _3MB_Trace trace) {
        this.traces.put(uuid, trace);
    }

    public _3MB_Trace get(UUID uuid) {
        return this.traces.get(uuid);
    }

    public void remove(UUID uuid) {
        this.traces.remove(uuid);
    }
}
