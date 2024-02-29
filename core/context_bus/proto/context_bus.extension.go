package proto

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"

	"errors"
	"fmt"
)

// Be careful of how ProtoBuffer deals with variable-size data structures such as map and slice
// We need to distinguish if they are empty or null.
// For example, an empty map (len = 0) will not be encoded during Marshal(),
// and thus we will get a null value for the map (ptr = nil) after Unmarshal().
// Instances of ProtoBuffer messages come from both generated code (ensures non-null values) and Unmarshal(),
// Make sure to check null values before taking operations.
// Most methods are designed for js-style method chaining

func (m *PrerequisiteSnapshot) Clone() *PrerequisiteSnapshot {
	if m == nil {
		return nil
	} else if m.Value == nil {
		return &PrerequisiteSnapshot{}
	}

	n := &PrerequisiteSnapshot{
		Value: make([]int64, len(m.Value)),
		Acc:   m.Acc,
	}

	copy(n.Value, m.Value)

	return n
}

func (m *PrerequisiteSnapshot) MergeOffset(src *PrerequisiteSnapshot) {
	for i := 0; i < len(m.Value); i++ {
		m.Value[i] += src.Value[i]
	}
}

func (m *PrerequisiteSnapshots) Clone() *PrerequisiteSnapshots {
	if m == nil {
		return nil
	} else if m.Snapshots == nil {
		return &PrerequisiteSnapshots{}
	}

	n := &PrerequisiteSnapshots{Snapshots: make(map[string]*PrerequisiteSnapshot, len(m.Snapshots))}
	for name, snapshot := range m.Snapshots {
		n.Snapshots[name] = snapshot.Clone()
	}

	return n
}

func (m *PrerequisiteSnapshots) GetPrerequisiteSnapshot(name string) *PrerequisiteSnapshot {
	if m == nil || m.Snapshots == nil {
		return nil
	}

	return m.Snapshots[name]
}

func (m *PrerequisiteSnapshots) MergeOffset(src *PrerequisiteSnapshots) {
	for name, dstS := range m.Snapshots {
		dstS.MergeOffset(src.Snapshots[name])
	}
}

func (m *Attributes) Clone() *Attributes {
	if m == nil {
		return nil
	} else if m.Attrs == nil {
		return &Attributes{}
	}

	n := &Attributes{Attrs: make(map[string]*AttributeValue, len(m.Attrs))}
	for key, val := range m.Attrs {
		n.Attrs[key] = val.Clone()
	}

	return n
}

func (m *Attributes) GetValue(path []string) (string, error) {
	attrs := m.Attrs

	for i, name := range path {
		if val, ok := attrs[name]; ok {
			if val.Type == AttributeValueType_AttributeValueStr {
				if i == len(path)-1 {
					return val.Str, nil
				}

				return "", errors.New("invalid Path for a string value")
			} else if val.Type == AttributeValueType_AttributeValueAttr {
				if i == len(path)-1 { // todo: not recommended
					return fmt.Sprintf("%+v", val), nil
				}

				return val.Struct.GetValue(path[1:])
			} else {
				return "", errors.New("invalid AttributeValueType")
			}
		} else {
			break
		}
	}

	return "", errors.New("value not found")
}

func (m *Attributes) Merge(attrs *Attributes) *Attributes {
	if attrs != nil {
		if m.Attrs == nil {
			m.Attrs = attrs.Attrs
		} else {
			for key, value := range attrs.Attrs {
				if val, ok := m.Attrs[key]; !ok {
					m.Attrs[key] = value
				} else if val.Type == AttributeValueType_AttributeValueAttr {
					val.Struct.Merge(value.Struct)
				}
			}
		}
	}

	return m
}

// SetString sets a string {value} for the {key},
// return {m} (self) for method chaining.
func (m *Attributes) SetString(key string, value string) *Attributes {
	m.Attrs[key] = &AttributeValue{Type: AttributeValueType_AttributeValueStr, Str: value}

	return m
}

// WithAttributes merges {attrs} for the value of given {key},
// returns the merged value of {key} for method chaining.
func (m *Attributes) WithAttributes(key string, attrs *Attributes) *Attributes {
	value := &AttributeValue{Type: AttributeValueType_AttributeValueAttr, Struct: attrs}
	if attrs == nil { // check null pointer for ProtoBuffer Unmarshal()
		value.Struct = &Attributes{Attrs: map[string]*AttributeValue{}}
	}

	if m.Attrs == nil { // check null pointer for ProtoBuffer Unmarshal()
		m.Attrs = map[string]*AttributeValue{}
	} else if val, ok := m.Attrs[key]; ok {
		value.Merge(val)
	}

	m.Attrs[key] = value

	return value.Struct
}

func (m *Attributes) SetAttributes(key string, value *Attributes) *Attributes {

	return m
}

func (m *Attributes) GetString(key string) (string, bool) {
	val, ok := m.Attrs[key]
	if ok && val.Type == AttributeValueType_AttributeValueStr {
		return val.Str, true
	}

	return "", false
}

func (m *EventWhen) Merge(when *EventWhen) *EventWhen {
	if when != nil {
		if m.Time == 0 {
			m.Time = when.Time
		}
	}

	return m
}

func (m *AttributeValue) Clone() *AttributeValue {
	if m.Type == AttributeValueType_AttributeValueAttr {
		return &AttributeValue{Type: m.Type, Struct: m.Struct.Clone()}
	}

	return &AttributeValue{Type: m.Type, Str: m.Str}
}

func (m *AttributeValue) Merge(value *AttributeValue) *AttributeValue {
	if value != nil {
		if m.Type == AttributeValueType_AttributeValueAttr && value.Type == AttributeValueType_AttributeValueAttr {
			m.Struct.Merge(value.Struct)
		}
	}

	return m
}

func (m *EventWhere) Merge(where *EventWhere) *EventWhere {
	if where != nil {
		if m.Attrs == nil {
			m.Attrs = where.Attrs
		} else {
			m.Attrs.Merge(where.Attrs)
		}

		if m.Stacktrace == "" {
			m.Stacktrace = where.Stacktrace
		}
	}

	return m
}

func (m *EventRecorder) Merge(recorder *EventRecorder) *EventRecorder {
	if recorder != nil {
		if m.Type == EventRecorderType_EventRecorderType_ {
			m.Type = recorder.Type
		}
		if m.Name == "" {
			m.Name = recorder.Name
		}
	}

	return m
}

func (m *EventMessage) GetValue(path []string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("invalid Path length for EventMessage")
	} else if len(path) == 1 && path[0] == "__message__" {
		return m.Message, nil
	}

	return m.Attrs.GetValue(path)
}

func (m *EventMessage) SetString(path []string, value string) *EventMessage {
	//m.Attrs.SetString(key, value)

	return m
}

func (m *EventMessage) SetAttributes(attrs *Attributes) *EventMessage {
	m.Attrs = attrs

	return m
}

func (m *EventMessage) GetAttributes() *Attributes {
	if m.Attrs == nil {
		m.Attrs = &Attributes{Attrs: map[string]*AttributeValue{}}
	}

	return m.Attrs
}

// todo: move to Attirbutes implementation
func (m *EventMessage) WithAttributes(attrs *Attributes) *Attributes {
	if attrs == nil {
		attrs = &Attributes{Attrs: map[string]*AttributeValue{}}
	}

	if m.Attrs == nil {
		m.Attrs = attrs
	} else {
		m.Attrs.Merge(attrs)
	}

	return m.Attrs
}

func (m *EventMessage) SetMessage(msg string) *EventMessage {
	m.Message = msg

	return m
}

func (m *EventMessage) SetPaths(paths []*Path) *EventMessage {
	m.Paths = paths

	return m
}

func (m *EventMessage) Merge(msg *EventMessage) *EventMessage {
	if msg != nil {
		if m.Attrs == nil {
			m.Attrs = msg.Attrs
		} else {
			m.Attrs.Merge(msg.Attrs)
		}

		if m.Message == "" {
			m.Message = msg.Message
		}
	}

	return m
}

func (m *LibrariesMessage) GetValue(path []string) (string, error) {
	if m == nil {
		return "", errors.New("no libraries")
	} else if len(m.Libraries) == 0 {
		return "", errors.New("empty libraries")
	} else if len(path) == 0 {
		return "", errors.New("invalid Path len for LibrariesMessage")
	}

	if lib, ok := m.Libraries[path[0]]; ok {
		return lib.GetValue(path[1:])
	}

	return "", errors.New("library not found")
}

func (m *LibrariesMessage) Merge(libs *LibrariesMessage) *LibrariesMessage {
	if libs != nil {
		if m.Libraries == nil {
			m.Libraries = libs.Libraries
		} else {
			for key, value := range libs.Libraries {
				if val, ok := m.Libraries[key]; !ok {
					m.Libraries[key] = value
				} else {
					val.Merge(value)
				}
			}
		}
	}

	return m
}

func (m *EventWhat) GetValue(path *Path) (string, error) {
	if path.Type == PathType_Application {
		return m.Application.GetValue(path.Path)
	}

	return m.Libraries.GetValue(path.Path)
}

func (m *EventWhat) Merge(what *EventWhat) *EventWhat {
	if what != nil {
		if m.Application == nil {
			m.Application = what.Application
		} else {
			m.Application.Merge(what.Application)
		}

		if m.Libraries == nil {
			m.Libraries = what.Libraries
		} else {
			m.Libraries.Merge(what.Libraries)
		}
	}

	return m
}

func (m *EventWhat) WithApplication(msg *EventMessage) *EventMessage {
	if m.Application == nil {
		if msg == nil {
			m.Application = &EventMessage{}
		} else {
			m.Application = msg
		}
	} else {
		m.Application.Merge(msg)
	}

	return m.Application
}

func (m *EventWhat) WithLibrary(key string, msg *EventMessage) *EventMessage {
	if msg == nil {
		msg = &EventMessage{}
	}

	if m.Libraries == nil {
		m.Libraries = &LibrariesMessage{
			Libraries: map[string]*EventMessage{key: msg},
		}
	} else if val, ok := m.Libraries.Libraries[key]; ok {
		msg.Merge(val)
	}

	m.Libraries.Libraries[key] = msg

	return msg
}

func (m *EventRepresentation) WithWhen(when *EventWhen) *EventWhen {
	if m.When == nil {
		if when == nil {
			m.When = &EventWhen{}
		} else {
			m.When = when
		}
	} else {
		m.When.Merge(when)
	}

	return m.When
}

func (m *EventRepresentation) WithWhere(where *EventWhere) *EventWhere {
	if m.When == nil {
		if where == nil {
			m.Where = &EventWhere{}
		} else {
			m.Where = where
		}
	} else {
		m.Where.Merge(where)
	}

	return m.Where
}

func (m *EventRepresentation) WithRecorder(recorder *EventRecorder) *EventRecorder {
	if m.Recorder == nil {
		if recorder == nil {
			m.Recorder = &EventRecorder{}
		} else {
			m.Recorder = recorder
		}
	} else {
		m.Recorder.Merge(recorder)
	}

	return m.Recorder
}

func (m *EventRepresentation) WithWhat(what *EventWhat) *EventWhat {
	if m.What == nil {
		if what == nil {
			m.What = &EventWhat{}
		} else {
			m.What = what
		}
	} else { // merge
		m.What.Merge(what)
	}

	return m.What
}

func (m *EventData) GetPreviousEventData(name string) *EventData {
	prev := m.PrevEventData
	if prev == nil {
		// todo not found
		fmt.Println("todo not found")
	}
	found := prev.Event.Recorder.Name == name
	for !found && prev != nil {
		if prev.Event.Recorder.Name == name {
			found = true
			break
		}
		prev = prev.PrevEventData
	}

	if found {
		return prev
	}

	return nil
}

func (m *PrometheusOpts) ToPrometheusCounterOpts() (prometheus.CounterOpts, []string) {
	return prometheus.CounterOpts{
		Namespace:   m.Namespace,
		Subsystem:   m.Subsystem,
		Name:        m.Name,
		Help:        m.Help,
		ConstLabels: m.ConstLabels,
	}, m.LabelNames
}

func (m *PrometheusOpts) ToPrometheusGaugeOpts() (prometheus.GaugeOpts, []string) {
	return prometheus.GaugeOpts{
		Namespace:   m.Namespace,
		Subsystem:   m.Subsystem,
		Name:        m.Name,
		Help:        m.Help,
		ConstLabels: m.ConstLabels,
	}, m.LabelNames
}

func (m *PrometheusHistogramOpts) ToPrometheus() (prometheus.HistogramOpts, []string) {
	return prometheus.HistogramOpts{
		Namespace:   m.Namespace,
		Subsystem:   m.Subsystem,
		Name:        m.Name,
		Help:        m.Help,
		ConstLabels: m.ConstLabels,
		Buckets:     m.Buckets,
	}, m.LabelNames
}

func (m *PrometheusSummaryOpts) ToPrometheus() (prometheus.SummaryOpts, []string) {
	objs := map[float64]float64{}
	for _, obj := range m.Objectives {
		objs[obj.Key] = obj.Value
	}

	return prometheus.SummaryOpts{
		Namespace:   m.Namespace,
		Subsystem:   m.Subsystem,
		Name:        m.Name,
		Help:        m.Help,
		ConstLabels: m.ConstLabels,
		Objectives:  objs,
		MaxAge:      time.Duration(m.MaxAge),
		AgeBuckets:  m.AgeBuckets,
		BufCap:      m.BufCap,
	}, m.LabelNames
}
