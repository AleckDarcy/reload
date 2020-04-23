package io.grpc.tracer;

import java.lang.Long;
import java.lang.String;
import java.util.concurrent.locks.ReentrantReadWriteLock;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class Store {
    public static String uuid = UUID.randomUUID().toString();

    private static ReentrantReadWriteLock lock = new ReentrantReadWriteLock();
    private static Map<Long, Message.Trace> traces = new HashMap<Long, Message.Trace>();

    public static void SetTrace(long threadID, Message.Trace t) {
        lock.writeLock().lock();
        traces.put(threadID, t);
        lock.writeLock().unlock();
    }

    public static Message.Trace GetTrace(long threadID) {
        Message.Trace t = null;
        lock.readLock().lock();
        if (traces.containsKey(threadID)) {
            t = traces.get(threadID);
        }
        lock.readLock().unlock();

        return t;
    }

    public static Message.Trace FetchTrace(long threadID) {
        Message.Trace t = null;
        lock.readLock().lock();
        if (traces.containsKey(threadID)) {
            t = traces.get(threadID);
            traces.remove(threadID);
        }
        lock.readLock().unlock();

        return t;
    }

//    full version
//    private static Map<Integer, Traces> traces = new HashMap<Integer, Traces>();
//    public static class ContextMeta {
//        private int traceID;
//        private String uuid;
//
//        public ContextMeta(int traceID, String uuid) {
//            this.traceID = traceID;
//            this.uuid = uuid;
//        }
//    }
//
//    public static class Traces {
//        private Map<String, Message.Trace> traces;
//
//        public Traces() {
//
//        }
//    }
//
//    public static boolean CheckByContextMeta(ContextMeta meta) {
//        boolean ok = false;
//
//        lock.readLock().lock();
//        if (traces.containsKey(meta.traceID)) {
//            Traces ts = traces.get(meta.traceID);
//            if (ts.traces.containsKey(meta.uuid)) {
//                ok = true;
//            }
//        }
//        lock.readLock().unlock();
//
//        return ok;
//    }
//
//    public static Message.Trace GetByContextMeta(ContextMeta meta) {
//        Message.Trace t = null;
//
//        lock.readLock().lock();
//        if (traces.containsKey(meta.traceID)) {
//            Traces ts = traces.get(meta.traceID);
//            if (ts.traces.containsKey(meta.uuid)) {
//                t = copyTrace(ts.traces.get(meta.uuid));
//            }
//        }
//        lock.readLock().unlock();
//
//        return t;
//    }
//
//    public static void SetByContextMeta(ContextMeta meta, Message.Trace trace) {
//        lock.writeLock().lock();
//        if (traces.containsKey(meta.traceID)) {
//            Traces ts = traces.get(meta.traceID);
////            if (ts.traces.containsKey(meta.uuid)) {
////                Message.Trace t = ts.traces.get(meta.uuid);
////                mergeTrace(t, trace);
////            } else {
//                ts.traces.put(meta.uuid, trace);
////            }
//        } else {
//            Traces ts = new Traces();
//            final String uuid = meta.uuid;
//            final Message.Trace t = trace;
//            ts.traces = new HashMap<String, Message.Trace>() {
//                {
//                    put(uuid, t);
//                }
//            };
//
//            traces.put(meta.traceID, ts);
//        }
//        lock.writeLock().unlock();
//    }
//
//    // todo UpdateFunctionByContextMeta
//
//    public static void DeleteByContextMeta(ContextMeta meta) {
//
//    }
//
//    private static Message.Trace copyTrace(Message.Trace oldTrace) {
//        // todo
//
//        return oldTrace;
//    }
//
////    private static void mergeTrace(Message.Trace dst, Message.Trace src) {
////
////    }
}