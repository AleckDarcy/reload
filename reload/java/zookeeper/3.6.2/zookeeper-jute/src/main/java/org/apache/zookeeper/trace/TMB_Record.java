package org.apache.zookeeper.trace;

import org.apache.jute.BinaryInputArchive;
import org.apache.jute.BinaryOutputArchive;
import org.apache.jute.Record;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * helper class for org.apache.jute.Record
 */
public class TMB_Record {
    protected static final List<TMB_Event> EMPTY_EVENTS = new ArrayList<>(0);
    protected static final List<TMB_TFI> EMPTY_TFIS = new ArrayList<>(0);

    // serialize with extra bytes
    public static byte[] serialize(Record record, byte[] prefix, byte[] suffix) throws IOException {
        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        BinaryOutputArchive bos = BinaryOutputArchive.getArchive(baos);
        if (prefix != null) {
            baos.write(prefix);
        }
        record.serialize(bos, "");
        if (suffix != null) {
            baos.write(suffix);
        }
        baos.close();

        return baos.toByteArray();
    }

    public static byte[] serialize(Record record) throws IOException {
        return serialize(record, null, null);
    }

    public static void deserialize(Record record, ByteArrayInputStream in) throws IOException {
        BinaryInputArchive ia = BinaryInputArchive.getArchive(in);
        record.deserialize(ia, "");
    }

    public static String getString(Record record) {
        if (record == null) {
            return "null";
        }

        String str = record.toString();
        while (str.endsWith("\n")) {
            str = str.substring(0, str.length() - 1);
        }

        return str;
    }

    public static String getTraceJson(Record record) {
        if (record == null) {
            return "{}";
        }

        return record.getTrace().toJSON();
    }

    // TODO: a return a boolean value or the number of new events
    /**
     * Appends an events with eventType, messange Name and uuid.
     * More specifically, appends the first ever event associated with a uuid.
     * Make sure: 1) trace is valid (id != 0); 2) TODO: a check events is not empty?
     * @param procMeta
     * @param record
     * @param eventType
     * @param messageName
     * @param uuid
     * @return
     */
    private static Record appendEventUnsafe(TMB_Store.ProcessorMeta procMeta, Record record, int eventType, String messageName, String uuid) {
        TMB_Trace trace = record.getTrace();
        trace.addEvent(new TMB_Event(eventType, messageName, uuid, procMeta));

        return record;
    }

    /**
     * Appends new event with eventType, messageName, uuid.
     * @param procMeta
     * @param record
     * @param eventType
     * @param messageName
     * @param uuid
     * @return
     */
    public static Record appendEvent(TMB_Store.ProcessorMeta procMeta, Record record, int eventType, String messageName, String uuid) {
        List<TMB_Event> events = record.getTrace().getEvents();
        int eventSize = events.size();
        if (eventSize > 0) {
            return appendEventUnsafe(procMeta, record, eventType, messageName, uuid);
        }

        return record;
    }

    /**
     * Appends new event with eventType and messageName.
     * Other information (uuid) is derived from last event of the current trace from record.
     * @param procMeta
     * @param record
     * @param eventType
     * @param messageName
     * @return
     */
    public static Record appendEvent(TMB_Store.ProcessorMeta procMeta, Record record, int eventType, String messageName) {
        List<TMB_Event> events = record.getTrace().getEvents();
        int eventSize = events.size();
        if (eventSize > 0) {
            TMB_Event lastEvent = events.get(eventSize - 1);
            String uuid = lastEvent.getUuid();

            return appendEventUnsafe(procMeta, record, eventType, messageName, uuid);
        }

        return record;
    }

    /**
     * Appends new event with eventType.
     * Other information (messageName, uuid) are derived from last event of the current trace from record.
     * @param procMeta
     * @param record
     * @param eventType
     * @return
     */
    public static Record appendEvent(TMB_Store.ProcessorMeta procMeta, Record record, int eventType) {
        List<TMB_Event> events = record.getTrace().getEvents();
        int eventSize = events.size();
        if (eventSize > 0) {
            TMB_Event lastEvent = events.get(eventSize - 1);
            String uuid = lastEvent.getUuid();
            String messageName = lastEvent.getMessage_name();

            return appendEventUnsafe(procMeta, record, eventType, uuid, messageName);
        }

        return record;
    }
}
