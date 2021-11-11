package log

import (
	"encoding/json"
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

const (
	NormalLogger = iota
	DebugHelperLogger
	StubLogger
	CriticalPathLogger
	TraceLogger
	ErrorLogger
	count
)

type logger struct {
	type_ int
}

var conf = make([]bool, count)

func init() {
	for i := range conf {
		conf[i] = true
	}
}

func SetLogger(type_ uint, on bool) {
	if type_ < count {
		conf[type_] = on
	}
}

func SetLoggers(types []uint, on bool) {
	for _, type_ := range types {
		SetLogger(type_, on)
	}
}

var Logger = logger{NormalLogger}
var Debug = logger{DebugHelperLogger}
var Stub = logger{StubLogger}
var CriticalPath = logger{CriticalPathLogger}
var Trace = logger{TraceLogger}
var Error = logger{ErrorLogger}

func (l *logger) PrintlnWithStackTrace(skip int, format string, a ...interface{}) {
	if conf[l.type_] == false {
		return
	}

	if skip == 2 {
		l.PrintlnWithCaller(format, a...)
	} else {
		f2, _ := caller2(2)
		f, line := caller2(skip)

		a = append([]interface{}{skip}, a...)

		Printf("[3MileBeach] "+f2+" (skip %d: "+f+") "+format+line, a...)
	}
}

func (l *logger) PrintlnWithCaller(format string, a ...interface{}) {
	if conf[l.type_] == false {
		return
	}

	f, line := caller2(2)

	Printf("[3MileBeach] "+f+" "+format+line, a...)
}

func (l *logger) Printf(format string, a ...interface{}) {
	if conf[l.type_] == false {
		return
	}

	Printf("[3MileBeach] "+format, a...)
}

type stringer struct{}

var Stringer stringer

func (s *stringer) JSON(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%+v", value)
	}

	return string(bytes)
}
