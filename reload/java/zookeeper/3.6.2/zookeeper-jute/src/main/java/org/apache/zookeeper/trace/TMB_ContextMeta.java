package org.apache.zookeeper.trace;

import java.util.UUID;

public class TMB_ContextMeta {
    private long trace_id;
    private UUID uuid;
    private String message_name;

    public TMB_ContextMeta(long trace_id, UUID uuid, String message_name) {
        this.trace_id = trace_id;
        this.uuid = uuid;
        this.message_name = message_name;
    }

    public long getTrace_id() {
        return trace_id;
    }

    public UUID getUuid() {
        return uuid;
    }

    public String getMessage_name() {
        return message_name;
    }
}
