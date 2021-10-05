package log

import (
	"fmt"
	"runtime"
	"strings"
)

func caller(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	_ = pc

	return fmt.Sprintf(" %s:%d\n", file, line)
}

func Caller(skip int) (string, string) {
	pc, file, line, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc).Name()
	if id := strings.Index(f, "("); id != -1 {
		f = f[id:]
	} else {
		id = strings.LastIndex(f, ".")
		f = f[id+1:]
	}

	return fmt.Sprintf("func %s()", f), fmt.Sprintf(" %s:%d\n", file, line)
}

func Logf(format string, v ...interface{}) {
	fmt.Printf(format+caller(2), v...)
}

func Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
