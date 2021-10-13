package milebeach

import (
	"github.com/AleckDarcy/reload/core/log"
)

type logger struct{}

var Logger *logger

func (l *logger) PrintlnWithStackTrace(skip int, format string, a ...interface{}) {
	f, line := log.Caller(skip)

	log.Printf("[3MileBeach] "+f+" "+format+line, a...)
}

func (l *logger) PrintlnWithCaller(format string, a ...interface{}) {
	f, line := log.Caller(2)

	log.Printf("[3MileBeach] "+f+" "+format+line, a...)
}

func (l *logger) Printf(format string, a ...interface{}) {
	log.Printf("[3MileBeach] "+format, a...)
}