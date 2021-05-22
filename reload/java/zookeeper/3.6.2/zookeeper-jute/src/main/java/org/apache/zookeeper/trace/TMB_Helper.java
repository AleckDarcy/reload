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
import java.util.concurrent.atomic.AtomicInteger;

public class TMB_Helper {
    private static AtomicInteger traceId = new AtomicInteger(0);
    private static AtomicInteger uuid = new AtomicInteger(0);

    public static boolean printable = true;

    public static String UUID() {
        return String.format("%010d", uuid.addAndGet(1));
//        return UUID.randomUUID().toString();
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

    public static void printf(TMB_Store.QuorumMeta quorumMeta, String format, Object ... args) {
        printf(quorumMeta, 3, format, args);
    }

    public static void printf(TMB_Store.QuorumMeta quorumMeta, int depth, String format, Object ... args) {
        if (printable) {
            StackTraceElement[] traces = Thread.currentThread().getStackTrace();
            if (traces.length <= depth) { // protection
                depth = traces.length - 1;
            }
            StackTraceElement trace = traces[depth];
            System.out.printf("[3MileBeach] %s:%d [%d] [%s] %s", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), quorumMeta.getName(), String.format(format, args));
        }
    }

    public static void printf(TMB_Store.ProcessorMeta procMeta, String format, Object ... args) {
        printf(procMeta, 3, format, args);
    }

    public static void printf(TMB_Store.ProcessorMeta procMeta, int depth, String format, Object ... args) {
        if (printable) {
            StackTraceElement[] traces = Thread.currentThread().getStackTrace();
            if (traces.length <= depth) { // protection
                depth = traces.length - 1;
            }
            StackTraceElement trace = traces[depth];
            System.out.printf("[3MileBeach] %s:%d [%d] [%s:%s] %s", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), procMeta.getQuorumName(), procMeta.getName(), String.format(format, args));
        }
    }

    public static String getClassNameFromObject(Object o) {
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

    public static String getClassName(Class c) {
        return getClassNameFromName(c.getCanonicalName());
    }

    public static String getClassNameFromName(String name) {
        return name.substring(name.lastIndexOf('.') + 1);
    }

    public static long currentTimeNanos() {
        return System.currentTimeMillis() * 1000000L + System.nanoTime() % 1000000L;
    }

    // helper function for debugging and testing
    public static int newTraceId() {
        return traceId.addAndGet(1);
    }
}
