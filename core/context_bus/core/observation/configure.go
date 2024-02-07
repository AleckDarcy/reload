package observation

import (
	cb "github.com/AleckDarcy/reload/core/context_bus/proto"
)

// implementations

func (c *Configure) Do(er *cb.EventRepresentation) {
	(*LoggingConfigure)(c.Logging).Do(er)
	(*TracingConfigure)(c.Tracing).Do(er)
	for _, metric := range c.Metrics {
		(*MetricsConfigure)(metric).Do(er)
	}
}

func (c *TimestampConfigure) Do() {
	if c != nil {
		return
	}
}

func (c *StackTraceConfigure) Do() {

}

func (c *LoggingConfigure) Do(er *cb.EventRepresentation) interface{} {
	if c == nil {
		return nil
	}

	(*TimestampConfigure)(c.Timestamp).Do()
	(*StackTraceConfigure)(c.Stacktrace).Do()

	tags := map[string]string{}

	for _, path := range c.Attrs {
		str, err := er.What.GetValue(path.Path)

		if err == nil {
			tags[path.Name] = str
		}
	}

	return tags
}

func (c *TracingConfigure) Do(er *cb.EventRepresentation) {
	if c == nil {
		return
	}
}

func (c *MetricsConfigure) Do(er *cb.EventRepresentation) {
	if c == nil {
		return
	}
}
