package proto

import (
	"errors"
	"fmt"
)

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

				return "", errors.New("TODO")
			} else if val.Type == AttributeValueType_AttributeValueAttr {
				if i == len(path)-1 { // todo: not recommended
					return fmt.Sprintf("%+v", val), nil
				}

				return val.Struct.GetValue(path[1:])
			} else {
				return "", errors.New("TODO")
			}
		} else {
			return "", errors.New("TODO")
		}
	}

	return "", errors.New("TODO")
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

func (m *Attributes) SetString(key string, value string) *Attributes {
	// todo init
	m.Attrs[key] = &AttributeValue{Type: AttributeValueType_AttributeValueStr, Str: value}

	return m
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

func (m *AttributeValue) Merge (value *AttributeValue) *AttributeValue {
	if value != nil {
		if m.Type == AttributeValueType_AttributeValueAttr && value.Type == AttributeValueType_AttributeValueAttr{
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

func (m *EventMessage) Init() {
	if m.Attrs == nil {
		m.Attrs = &Attributes{
			Attrs: map[string]*AttributeValue{},
		}
	} else if m.Attrs.Attrs == nil {
		m.Attrs.Attrs = map[string]*AttributeValue{}
	}
}

func (m *EventMessage) GetValue(path []string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("TODO")
	} else if len(path) == 1 && path[0] == "__message__" {
		return m.Message, nil
	}

	return m.Attrs.GetValue(path)
}

func (m *EventMessage) SetString(key string, value string) *EventMessage {
	m.Init()
	m.Attrs.SetString(key, value)

	return m
}

func (m *EventMessage) SetAttributes(key string, attrs *Attributes) *EventMessage {
	m.Init()
	m.Attrs.Attrs[key] = &AttributeValue{Type: AttributeValueType_AttributeValueAttr, Struct: attrs}

	return m
}

func (m *EventMessage) WithAttributes(key string, attrs *Attributes) *Attributes {
	value := &AttributeValue{Type: AttributeValueType_AttributeValueAttr, Struct: attrs}
	if attrs == nil {
		value.Struct = &Attributes{map[string]*AttributeValue{}}
	}

	if m.Attrs == nil {
		m.Attrs = &Attributes{
			Attrs: map[string]*AttributeValue{},
		}
	} else if val, ok := m.Attrs.Attrs[key]; ok {
		value.Merge(val)
	}

	m.Attrs.Attrs[key] = value

	return value.Struct
}

func (m *EventMessage) SetMessage(msg string) *EventMessage {
	m.Message = msg

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
	if len(path) == 0 {
		return "", errors.New("TODO")
	}

	if lib, ok := m.Libraries[path[0]]; ok {
		return lib.GetValue(path[1:])
	}

	return "", errors.New("TODO")
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
