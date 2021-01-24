package org.apache.zookeeper.trace;

import org.apache.jute.Record;

import java.util.*;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class TMB_Store {
    class ClientPluginTrace {

    }

    class QuorumTrace {

    }

    private static Map<Long, ClientPluginTrace> clientPluginTraces = new HashMap<>();
    private static Map<Long, QuorumTrace> quorumTrace = new HashMap<>();

    private static Map<Long, TMB_Trace> thread_traces = new HashMap<>();
    private static Map<Long, TMB_Trace> server_traces = new HashMap<>();
    private static ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    public static void clearServerTraces() {
        lock.writeLock().lock();
        server_traces.clear();
        lock.writeLock().unlock();
    }

    public static void serverSetTrace(long serverId, TMB_Trace trace) {
        if (trace.getId() == 0) {
            return;
        }

        lock.writeLock().lock();
        server_traces.put(serverId, trace);
        lock.writeLock().unlock();
    }

    public static void serverMergeTrace(long serverId, TMB_Trace trace) {
        if (trace == null || trace.getId() == 0) {
            return;
        }

        List<TMB_Event> events = trace.getEvents();
        int reqEvent = (int) trace.getReqEvent();
        if (events.size() <= reqEvent) {
            return;
        }

        events = events.subList(reqEvent, events.size());

        lock.writeLock().lock();
        TMB_Trace trace_ = server_traces.get(serverId);
        if (trace_ == null || trace_.getId() != trace.getId()) {
            lock.writeLock().unlock();

            return;
        }

        List<TMB_Event> events_ = trace_.getEvents();
        for (TMB_Event event: events) {
            boolean found = false;
            for (TMB_Event event_: events_) {
                found = event_.equals(event);
                if (found) {
                    break;
                }
            }
            if (!found) {
                events_.add(event);
            }
        }
        trace_.setEvents(events_);

        lock.writeLock().unlock();
    }

    public static TMB_Trace serverGetThread(long serverId) {
        lock.readLock().lock();
        TMB_Trace trace = server_traces.get(serverId);
        lock.readLock().unlock();

        return trace;
    }

    public static void callerAppendEventsByThreadIdUnsafe(long threadId, TMB_Trace trace) {
        TMB_Trace trace_ = thread_traces.get(threadId);

        if (trace_ == null) {
            thread_traces.put(threadId, trace);
        } else {
//            TMB_Helper.println("before appending, length: " + trace_.getEvents().size() + " + " + trace.getEvents().size());

            List<TMB_Event> events_ = trace_.getEvents();
            List<TMB_Event> events = trace.getEvents();

            events_.addAll(events);
            trace_.setEvents(events_);

//            TMB_Helper.println("after appending, length: " + trace_.getEvents().size());
        }
    }

    public static TMB_Trace getByThreadId(long threadId) {
        lock.readLock().lock();
        TMB_Trace trace = thread_traces.get(threadId);
        lock.readLock().unlock();

        return trace;
    }

    // TODO: 3MileBeach temp function
    public static Map<Long, TMB_Trace> getAllThreads() {
        Map<Long, TMB_Trace> result = new HashMap<>();
        lock.readLock().lock();
        thread_traces.forEach(result::put);
        lock.readLock().unlock();

        return result;
    }

    public static void removeByThreadId(long threadId) {
        lock.writeLock().lock();
        thread_traces.remove(threadId);
        lock.writeLock().unlock();
    }

    /**
     * Server receives request from client or upstream server
     * RequestProcessor -> calleeInbound -> process request -> calleeOutbound -> send response
     */
    // to be called after ByteBufferInputStream.byteBuffer2Record(buffer, record)
    public static void calleeInbound(String service, Record request) {
        TMB_Trace trace = request.getTrace();
        if (trace == null || trace.getId() == 0) {
            return;
        }

        List<TMB_Event> events = trace.getEvents();
        if (events.size() != 1) {
            TMB_Helper.println("callee inbound receives request with invalid trace: " + TMB_Helper.getClassName(request) + "(" + TMB_Helper.getString(request) + ")");

            return;
        }

        TMB_Helper.println("callee inbound receives request: " + TMB_Helper.getClassName(request) + "(" + TMB_Helper.getString(request) + ")");

        long threadId = Thread.currentThread().getId();
        TMB_Event preEvent = events.get(0);
        TMB_Event event = new TMB_Event(TMB_Event.RECORD_RECV, TMB_Helper.currentTimeNanos(), preEvent.getMessage_name(), preEvent.getUuid(), service);
        events.set(0, event);
        trace.setEvents(events);

        lock.writeLock().lock();
        thread_traces.put(threadId, trace);
        lock.writeLock().unlock();
    }

    public static void calleeOutbound(String service, Record response) {
        long threadId = Thread.currentThread().getId();

        lock.writeLock().lock();
        TMB_Trace trace = thread_traces.get(threadId);
        thread_traces.remove(threadId);
        lock.writeLock().unlock();

        if (trace == null) {
            TMB_Helper.println("callee outbound ejects response without trace: " + TMB_Helper.getClassName(response) + "(" + TMB_Helper.getString(response) + ")");

            return;
        }

        TMB_Event preEvent = trace.getEvents().get(0);
        TMB_Event event = new TMB_Event(TMB_Event.RECORD_SEND, TMB_Helper.currentTimeNanos(), TMB_Helper.getClassName(response), preEvent.getUuid(), service);

        trace.addEvent(event);
        response.setTrace(trace);

        TMB_Helper.println("callee outbound ejects response: " + TMB_Helper.getClassName(response) + "(" + TMB_Helper.getString(response) + ")");
    }

    /**
     * TODO: Client (Server when sending request to downstream server)
     * submitRequest -> generate request -> callerOutbound -> network -> callerInbound -> process response
     */
    public static void callerOutbound(String service, Record request) {
        TMB_Trace trace = request.getTrace();
        // stub, should be called only once per client-level request
        // TODO: let client generate trace_id
        if (trace.getId() == 0) {
            long id = TMB_Helper.newTraceId();
            trace.setId(id);
            trace.setEvents(new ArrayList<>());
            TMB_Helper.println("stub trace with id:" + id);
            new Exception().printStackTrace();
        }

        if (trace.getId() != 0) {
            long threadId = Thread.currentThread().getId();
            String requestName = TMB_Helper.getClassName(request);
            String uuid = UUID.randomUUID().toString();
            TMB_Event event = new TMB_Event(TMB_Event.RECORD_SEND, TMB_Helper.currentTimeNanos(), requestName, uuid, service);

            List<TMB_Event> events = trace.getEvents();
            events.add(event);
            trace.setEvents(events);

            lock.writeLock().lock();
            TMB_Store.callerAppendEventsByThreadIdUnsafe(threadId, trace);
            lock.writeLock().unlock();
        }

        TMB_Helper.println("caller outbound ejects request: " + TMB_Helper.getClassName(request) + "(" + TMB_Helper.getString(request) + ")");
    }

    public static void callerInbound(String service, Record response) {
        TMB_Trace trace = response.getTrace();
        if (trace.getId() == 0) {
            return;
        }

        String responseName = TMB_Helper.getClassName(response);
        List<TMB_Event> events = trace.getEvents();
        if (events.size() == 0) {
            TMB_Helper.println("caller inbound receives trace without events: " + responseName + "(" + TMB_Helper.getString(response) + ")");

            return;
        }

        long threadId = Thread.currentThread().getId();
        String uuid = events.get(0).getUuid();
        TMB_Event event = new TMB_Event(TMB_Event.RECORD_RECV, TMB_Helper.currentTimeNanos(), responseName, uuid, service);
        trace.addEvent(event);

        lock.writeLock().lock();
        TMB_Store.callerAppendEventsByThreadIdUnsafe(threadId, trace);
        lock.writeLock().unlock();

        TMB_Helper.println("caller inbound receives response: " + responseName + "(" + TMB_Helper.getString(response) + ")");
    }
}
