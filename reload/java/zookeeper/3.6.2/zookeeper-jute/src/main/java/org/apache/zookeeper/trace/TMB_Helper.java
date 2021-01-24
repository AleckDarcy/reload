package org.apache.zookeeper.trace;

import org.apache.jute.BinaryInputArchive;
import org.apache.jute.BinaryOutputArchive;
import org.apache.jute.Record;
import org.apache.zookeeper.proto.NullPointerRequest;
import org.apache.zookeeper.proto.NullPointerResponse;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.lang.String;
import java.lang.Thread;
import java.util.concurrent.atomic.AtomicInteger;

public class TMB_Helper {
    private static AtomicInteger traceId = new AtomicInteger(0);

    public static void println(String x) {
        StackTraceElement trace = Thread.currentThread().getStackTrace()[2];

        System.out.printf("[3MileBeach] %s:%d [%d] %s\n", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), x);
    }

    public static void printf(String format, Object ... args) {
        StackTraceElement trace = Thread.currentThread().getStackTrace()[2];

        System.out.printf("[3MileBeach] %s:%d [%d] %s", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), String.format(format, args));
    }

    public static String getClassName(Object o) {
        if (o == null) {
            return "NullPointer";
        }

        if (o instanceof NullPointerRequest) {
            return ((NullPointerRequest) o).getRequestName() + "(*)";
        } else if (o instanceof NullPointerResponse) {
            return ((NullPointerResponse) o).getRequestName() + "(*)";
        }

        String name = o.getClass().getCanonicalName();

        return name.substring(name.lastIndexOf('.') + 1);
    }

    public static ByteArrayOutputStream serialize(Record record) throws IOException {
        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        BinaryOutputArchive bos = BinaryOutputArchive.getArchive(baos);
        record.serialize(bos, "");
        baos.close();

        return baos;
    }

    public static void deserialize(ByteArrayInputStream in, Record record) throws IOException {
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

        TMB_Trace trace = record.getTrace();
        if (trace == null) {
            return "{}";
        }

        return trace.toJSON();
    }

    public static long currentTimeNanos() {
        return System.currentTimeMillis() * 1000000L + System.nanoTime() % 1000000L;
    }

    // helper function for debugging and testing
    public static int newTraceId() {
        return traceId.addAndGet(1);
    }
}
