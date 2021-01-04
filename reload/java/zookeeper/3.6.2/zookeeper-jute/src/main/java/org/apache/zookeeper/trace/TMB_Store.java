package org.apache.zookeeper.trace;

import org.apache.jute.Record;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class TMB_Store {
    private static Map<Long, TMB_Trace> thread_traces = new HashMap<>();
    private static Map<UUID, TMB_ThreadTraces> request_traces = new HashMap<>();
    private static ReentrantReadWriteLock lock = new ReentrantReadWriteLock();

    /**
     * Server receives request from client or upstream server
     * RequestProcessor -> requestInbound -> process request -> requestOutbound -> send response
     */
    // to be called after ByteBufferInputStream.byteBuffer2Record(buffer, record)
    public static void requestInbound(Record request) {
        TMB_Trace trace = request.getTrace();
        if (trace == null || trace.getId() == 0) {
            return;
        }

        List<TMB_Event> events = trace.getEvents();
        if (events.size() == 0) {
            TMB_Helper.println("request inbound receives request with invalid trace: " + TMB_Helper.getClassName(request) + "(" + TMB_Helper.getString(request) + ")");

            return;
        }

        TMB_Helper.println("request inbound receives request: " + TMB_Helper.getClassName(request) + "(" + TMB_Helper.getString(request) + ")");

        try {
            long threadId = Thread.currentThread().getId();
            lock.writeLock().lock();

            TMB_Event preEvent = events.get(0);
            TMB_Event event = new TMB_Event(TMB_Event.RECORD_SEND, TMB_Helper.currentTimeNanos(), preEvent.getMessage_name(), preEvent.getUuid(), "TODO");

            events.add(event);
            trace.setEvents(events);

            thread_traces.put(threadId, trace);
        } finally {
            lock.writeLock().unlock();
        }
    }

    public static void requestOutbound(Record response) {
        try {
            long threadId = Thread.currentThread().getId();
            lock.writeLock().lock();

            TMB_Trace trace = thread_traces.get(threadId);
            if (trace == null) {
                TMB_Helper.println("request outbound ejects response without trace: " + TMB_Helper.getClassName(response) + "(" + TMB_Helper.getString(response) + ")");

                return;
            }
            thread_traces.remove(threadId);

            TMB_Event preEvent = trace.getEvents().get(0);
            TMB_Event event = new TMB_Event(TMB_Event.RECORD_SEND, TMB_Helper.currentTimeNanos(), TMB_Helper.getClassName(response), preEvent.getUuid(), "TODO");

            trace.addEvent(event);
            response.setTrace(trace);

            TMB_Helper.println("request outbound ejects response: " + TMB_Helper.getClassName(response) + "(" + TMB_Helper.getString(response) + ")");

        } finally {
            lock.writeLock().unlock();
        }
    }

    /**
     * TODO: Client (Server when sending request to downstream server)
     * submitRequest -> generate request -> responseOutbound -> network -> responseInbound -> process response
     */
    public static void responseOutbound(Record request) {
        try {
            int threadId = 0;
            lock.writeLock().lock();

            TMB_Helper.println("response outbound: " + TMB_Helper.getClassName(request) + "(" + TMB_Helper.getString(request) + ")");

        } finally {
            lock.writeLock().unlock();
        }
    }

    public static void responseInbound(Record response) {
        try {
            int threadId = 0;
            lock.writeLock().lock();

            TMB_Helper.println("response inbound: " + TMB_Helper.getClassName(response) + "(" + TMB_Helper.getString(response) + ")");

        } finally {
            lock.writeLock().unlock();
        }
    }
}
