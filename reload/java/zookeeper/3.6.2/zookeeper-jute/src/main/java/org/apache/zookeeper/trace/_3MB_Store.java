package org.apache.zookeeper.trace;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class _3MB_Store {
    private Map<Long, _3MB_ThreadTraces> thread_traces;
    private Map<UUID, _3MB_ThreadTraces> request_traces;
    private ReentrantReadWriteLock lock;

    public _3MB_Store() {
        this.thread_traces = new HashMap<Long, _3MB_ThreadTraces>();
        this.request_traces = new HashMap<UUID, _3MB_ThreadTraces>();
        this.lock = new ReentrantReadWriteLock();
    }


}
