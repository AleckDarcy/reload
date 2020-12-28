package org.apache.zookeeper.trace;

import java.lang.String;
import java.lang.Thread;

public class TMB_Helper {
    public static void println(String x) {
        StackTraceElement trace = Thread.currentThread().getStackTrace()[2];

        System.out.printf("[3MileBeach] %s:%d [%d] %s\n", trace.getFileName(), trace.getLineNumber(), Thread.currentThread().getId(), x);
    }

    public static String getClassName(Object o) {
        if (o == null) {
            return "NullPointer";
        }

        return o.getClass().getCanonicalName();
    }
}
