package org.apache.zookeeper.trace;

import java.lang.String;
import java.lang.Thread;

public class _3MB_Helper {
    public static void println(String x) {
        StackTraceElement trace = Thread.currentThread().getStackTrace()[2];

        System.out.println(trace.getFileName() + ":" + trace.getLineNumber() + " " + x);
    }

    public static String getClassName(Object o) {
        if (o == null) {
            return "NullPointer";
        }

        return o.getClass().getCanonicalName();
    }
}
