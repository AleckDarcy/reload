package proto

import (
	"errors"
	"fmt"
)

// Extension of interface isAttributeValue_Value to support Clone()
// todo: after re-generating protobuf code, delete the redeclared definition inside pb.go
type isAttributeValue_Value interface{
	isAttributeValue_Value()
	Clone() isAttributeValue_Value
}

func (m *AttributeValue) Clone() *AttributeValue {
	return &AttributeValue{Type: m.Type, Value: m.Value.Clone()}
}

func (m *AttributeValue_Str) Clone() isAttributeValue_Value {
	return m
}

func (m *AttributeValue_Struct) Clone() isAttributeValue_Value {
	return &AttributeValue_Struct{
		Struct: m.Struct.Clone(),
	}
}

func (m *AttributeValue_Struct) Merge(value *AttributeValue) *AttributeValue_Struct {
	if value != nil {
		if value.Type == AttributeValueType_AttributeValueAttr {
			if m.Struct == nil {
				m.Struct = value.Value.(*AttributeValue_Struct).Struct
			} else {
				m.Struct.Merge(value.Value.(*AttributeValue_Struct).Struct)
			}
		}
	}

	return m
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
				if i == len(path) - 1 {
					return val.Value.(*AttributeValue_Str).Str, nil
				}

				return "", errors.New("TODO")
			} else if val.Type == AttributeValueType_AttributeValueAttr {
				if i == len(path) - 1 {
					return fmt.Sprintf("%+v", val), errors.New("TODO")
				}

				return val.Value.(*AttributeValue_Struct).Struct.GetValue(path[1:])
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
					val.Value.(*AttributeValue_Struct).Merge(value)
				}
			}
		}
	}

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
		return "", errors.New("TODO")
	} else if len(path) == 1 && path[0] == "message" {
		return m.Message, nil
	}

	return m.Attrs.GetValue(path)
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

func (m *LibrariesMessage) WithLibrary(name string, msg *EventMessage) *EventMessage {
	if m.Libraries == nil {
		m.Libraries = map[string]*EventMessage{
			name: msg,
		}
	} else {
		if val, ok := m.Libraries[name]; !ok {
			m.Libraries[name] = msg
		} else {
			val.Merge(msg)
		}
	}

	return m.Libraries[name]
}

func (m *EventWhat) GetValue(path []string) (string, error) {
	if len(path) <= 1 {
		return "", errors.New("TODO")
	}

	if path[0] == "__application__" {
		return m.Application.GetValue(path[1:])
	}

	return m.Libraries.GetValue(path[1:])
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

func (m *EventWhat) WithLibraries(libs *LibrariesMessage) *LibrariesMessage {
	if m.Libraries == nil {
		if libs == nil {
			m.Libraries = &LibrariesMessage{}
		} else {
			m.Libraries = libs
		}
	} else {
		m.Libraries.Merge(libs)
	}

	return m.Libraries
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
