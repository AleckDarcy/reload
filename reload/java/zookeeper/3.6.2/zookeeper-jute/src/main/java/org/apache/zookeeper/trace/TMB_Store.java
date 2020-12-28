package org.apache.zookeeper.trace;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class TMB_Store {
    private Map<Long, TMB_ThreadTraces> thread_traces;
    private Map<UUID, TMB_ThreadTraces> request_traces;
    private ReentrantReadWriteLock lock;

    public TMB_Store() {
        this.thread_traces = new HashMap<Long, TMB_ThreadTraces>();
        this.request_traces = new HashMap<UUID, TMB_ThreadTraces>();
        this.lock = new ReentrantReadWriteLock();
    }


}
