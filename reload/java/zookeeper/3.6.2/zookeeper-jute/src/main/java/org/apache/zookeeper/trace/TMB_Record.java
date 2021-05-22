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

    public static void addEvent(TMB_Store.ProcessorMeta procMeta, Record record, int eventType, String messageName, String uuid) {
        record.getTrace().addEvent(procMeta, eventType, messageName, uuid);
    }

    public static void addEvent(TMB_Store.ProcessorMeta procMeta, Record record, int eventType, String messageName) {
        record.getTrace().addEvent(procMeta, eventType, messageName);

    }

    public static void addEvent(TMB_Store.ProcessorMeta procMeta, Record record, int eventType) {
        record.getTrace().addEvent(procMeta, eventType);

    }
}
