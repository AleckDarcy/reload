package injector

import (
	"fmt"
	"runtime"
)

func caller(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	_ = pc

	return fmt.Sprintf(" %s:%d\n", file, line)
}

func Logf(format string, v ...interface{}) {
	fmt.Printf(format + caller(2), v...)
}
