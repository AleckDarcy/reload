package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// A list of type definitions from proto messages, in the order of their definition.

type AttributeConfiguration cb.AttributeConfigure
type TimestampConfigure cb.TimestampConfigure
type StackTraceConfigure cb.StackTraceConfigure
type LoggingConfigure cb.LoggingConfigure
type TracingConfigure cb.TracingConfigure
type MetricsConfigure cb.MetricsConfigure
type Configure cb.ObservationConfigure
