package org.apache.zookeeper.trace;

public class TMB_ClientPlugin {
    private final long threadID = Thread.currentThread().getId();

    /**
     *
     */
    public void TMB_Initialize() {

    }

    /**
     * Get and delete traces
     */
    public void TMB_Finalize() {
        TMB_Trace trace = TMB_Store.getByThreadId(threadID);
        TMB_Helper.println("Trace for thread " + threadID + ": " + trace.toJSON());

        TMB_Store.removeByThreadId(threadID);
    }
}
