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
import java.util.List;
import java.util.UUID;
import java.util.concurrent.atomic.AtomicInteger;

public class TMB_Helper {
    private static AtomicInteger traceId = new AtomicInteger(0);
    private static AtomicInteger uuid = new AtomicInteger(0);

    public static boolean printable = true;

    public static final int UUID_LEN = 10;
    public static final int UUID_SUF_LEN = 5;

    public static String UUID() {
        return String.format("%010d", uuid.addAndGet(1));
//        return UUID.randomUUID().toString();
    }

    public static void checkTFIs(TMB_Trace trace, String messageName) throws FaultInjectedException {
        List<TMB_TFI> tfis = trace.getTfis();
        for (TMB_TFI tfi: tfis) {
            if (tfi.getName().equals(messageName) && (tfi.getEvent_type() & TMB_Event.RECORD_SEND) != 0) {
                List<TMB_TFIMeta> metas = tfi.getAfter();
                boolean injected = true;
                if (metas != null && metas.size() > 0) {
                    for (TMB_TFIMeta meta : tfi.getAfter()) {
                        if (meta.getAlready() < meta.getTimes()) {
                            injected = false;
                            break;
                        }
                    }
                }

                if (injected) {
                    throw new FaultInjectedException(tfi.getType(), tfi.getDelay());
                }
            }
        }
    }

    public static void println(String x) {
        if (printable) {
            StackTraceElement trace = Thread.currentThread().getStackTrace()[2];
            System.out.printf("[3MileBeach] %s:%d [%d] %s\n", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), x);
        }
    }

    public static void printf(String format, Object ... args) {
        printf(3, format, args);
    }

    public static void printf(int depth, String format, Object ... args) {
        if (printable) {
            StackTraceElement[] traces = Thread.currentThread().getStackTrace();
            if (traces.length <= depth) { // protection
                depth = traces.length - 1;
            }
            StackTraceElement trace = traces[depth];
            System.out.printf("[3MileBeach] %s:%d [%d] %s", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), String.format(format, args));
        }
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

        return getClassNameFromName(o.getClass().getCanonicalName());
    }

    public static String getClassNameFromClass(Class c) {
        return getClassNameFromName(c.getCanonicalName());
    }

    public static String getClassNameFromName(String name) {
        return name.substring(name.lastIndexOf('.') + 1);
    }

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
