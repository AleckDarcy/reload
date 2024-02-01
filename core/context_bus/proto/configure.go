package proto

func (m *Configure) Do(er *EventRepresentation) {
	name := er.Recorder.Name

	cfg := m.Observations[name]
	if cfg == nil {
		return
	}

	cfg.Logging.Do(er)
	cfg.Tracing.Do(er)
	for _, metric := range cfg.Metrics {
		metric.Do(er)
	}
}

func (m *TimestampConfiguration) Do () {
	if m != nil {
		return
	}
}

func (m *StackTraceConfiguration) Do () {

}

func (m *LoggingConfigure) Do(er *EventRepresentation) interface{} {
	if m == nil {
		return nil
	}

	m.Timestamp.Do()
	m.Stacktrace.Do()

	tags := map[string]string{}

	for _, path := range m.Attrs {
		str, err := er.What.GetValue(path.Path)
		if err == nil {
			tags[path.Name] = str
		}
	}

	return tags
}

func (m *TracingConfigure) Do(er *EventRepresentation) {
	if m == nil {
		return
	}
}

func (m *MetricsConfigure) Do (er *EventRepresentation) {
	if m == nil {
		return
	}
}

func TTT()  {

}