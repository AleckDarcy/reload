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

func caller2(skip int) (string, string) {
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

type logger struct{}

var Logger *logger

func (l *logger) PrintlnWithStackTrace(skip int, format string, a ...interface{}) {
	f, line := caller2(skip)

	Printf("[3MileBeach] "+f+" "+format+line, a...)
}

func (l *logger) PrintlnWithCaller(format string, a ...interface{}) {
	f, line := caller2(2)

	Printf("[3MileBeach] "+f+" "+format+line, a...)
}

func (l *logger) Printf(format string, a ...interface{}) {
	Printf("[3MileBeach] "+format, a...)
}
