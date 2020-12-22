package org.apache.zookeeper.trace;

import java.lang.String;

public class Helper {
    public static String getClassName(Object o) {
        if (o == null) {
            return "NullPointer";
        }

        return o.getClass().getCanonicalName();
    }
}
