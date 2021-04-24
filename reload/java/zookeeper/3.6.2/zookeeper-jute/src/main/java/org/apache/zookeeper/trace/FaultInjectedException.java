package org.apache.zookeeper.trace;

public class FaultInjectedException extends Exception {
    private int type;
    private long delay;

    public FaultInjectedException(int type, long delay) {
        this.type = type;
        this.delay = delay;
    }

    public FaultInjectedException() {
        super();
    }

    public FaultInjectedException(String message) {
        super(message);
    }

    public FaultInjectedException(String message, Throwable cause) {
        super(message, cause);
    }

    public FaultInjectedException(Throwable cause) {
        super(cause);
    }

    public int GetType() {
        return type;
    }

    public long GetDelay() {
        return delay;
    }
}
