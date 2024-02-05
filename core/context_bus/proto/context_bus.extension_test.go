package proto

import (
	"reflect"
	"testing"
)

var attrs1 = &Attributes{
	Attrs: map[string]*AttributeValue{
		"key1": {
			Type: AttributeValueType_AttributeValueStr,
			Str:  "value1",
		},
		"key2": {
			Type: AttributeValueType_AttributeValueAttr,
			Struct: &Attributes{
				Attrs: map[string]*AttributeValue{
					"key21": {
						Type: AttributeValueType_AttributeValueStr,
						Str:  "value21",
					},
				},
			},
		},
	},
}

var attrs2 = &Attributes{
	Attrs: map[string]*AttributeValue{
		"key1": {
			Type: AttributeValueType_AttributeValueStr,
			Str:  "value2",
		},
		"key3": {
			Type: AttributeValueType_AttributeValueAttr,
			Struct: &Attributes{
				Attrs: map[string]*AttributeValue{
					"key31": {
						Type: AttributeValueType_AttributeValueStr,
						Str:  "value31",
					},
				},
			},
		},
	},
}

var attrs3 = &Attributes{
	Attrs: map[string]*AttributeValue{
		"key1": {
			Type: AttributeValueType_AttributeValueStr,
			Str:  "value1",
		},
		"key2": {
			Type: AttributeValueType_AttributeValueAttr,
			Struct: &Attributes{
				Attrs: map[string]*AttributeValue{
					"key21": {
						Type: AttributeValueType_AttributeValueStr,
						Str:  "value21",
					},
				},
			},
		},
		"key3": {
			Type: AttributeValueType_AttributeValueAttr,
			Struct: &Attributes{
				Attrs: map[string]*AttributeValue{
					"key31": {
						Type: AttributeValueType_AttributeValueStr,
						Str:  "value31",
					},
				},
			},
		},
	},
}

func TestAttributes_Merge(t *testing.T) {
	attrs1.Merge(attrs2)

	if !reflect.DeepEqual(attrs1, attrs3) {
		t.Error("fail")
	}
}

func BenchmarkAttributes_Merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		attrs1.Merge(attrs2)
	}
}

func TestAttributes_GetValue(t *testing.T) {
	t.Log(attrs3.GetValue([]string{"key2", "key21"}))
}

func BenchmarkAttributes_GetValue(b *testing.B) {
	key2 := attrs3.Attrs["key2"]
	key3 := key2.Struct
	for i := 0; i < b.N; i++ {
		//if key2.Type == AttributeValueType_AttributeValueAttr {
		//	key3 = key2.Value.(*AttributeValue_Struct).Struct
		//}
		attrs3.GetValue([]string{"key2", "key21"})
	}

	b.Log(key3)
}

var what = &EventWhat{
	Application: &EventMessage{
		Attrs: &Attributes{
			Attrs: map[string]*AttributeValue{
				"key1": {
					Type: AttributeValueType_AttributeValueStr,
					Str:  "value1",
				},
				"key2": {
					Type: AttributeValueType_AttributeValueAttr,
					Struct: &Attributes{
						Attrs: map[string]*AttributeValue{
							"key21": {
								Type: AttributeValueType_AttributeValueStr,
								Str:  "value21",
							},
						},
					},
				},
			},
		},
		Message: "application message",
	},
	Libraries: &LibrariesMessage{
		Libraries: map[string]*EventMessage{
			"lib1": {
				Attrs: &Attributes{
					Attrs: map[string]*AttributeValue{
						"key1": {
							Type: AttributeValueType_AttributeValueAttr,
							Struct: &Attributes{
								Attrs: map[string]*AttributeValue{
									"key11": {
										Type: AttributeValueType_AttributeValueStr,
										Str:  "value11",
									},
								},
							},
						},
						"key2": {
							Type: AttributeValueType_AttributeValueStr,
							Str:  "value2",
						},
					},
				},
				Message: "lib1 message",
			},
		},
	},
}

func TestEventWhat_GetValue(t *testing.T) {
	t.Log(what.GetValue(&Path{Type: PathType_Application, Path: []string{"__message__"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Application, Path: []string{"key1"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Library, Path: []string{"lib1", "__message__"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Library, Path: []string{"lib1", "key2"}}))

	// multi-layer, stringify struct (not recommended)
	t.Log(what.GetValue(&Path{Type: PathType_Application, Path: []string{"key2"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Application, Path: []string{"key2", "key21"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Library, Path: []string{"lib1", "key1"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Library, Path: []string{"lib1", "key1", "key11"}}))

	// value not found
	t.Log(what.GetValue(&Path{Type: PathType_Library, Path: []string{"lib2"}}))
	t.Log(what.GetValue(&Path{Type: PathType_Library, Path: []string{"lib2", "key1", "key11"}}))
}

func BenchmarkEventWhat_GetValue(b *testing.B) {
	path := &Path{Type: PathType_Application, Path: []string{"__message__"}}
	//path := &Path{Type: PathApplication, Path: []string{"key1"}}
	//path := &Path{Type: PathLibrary, Path: []string{"lib1", "__message__"}}
	//path := &Path{Type: PathLibrary, Path: []string{"lib1", "key2"}}

	for i := 0; i < b.N; i++ {
		what.GetValue(path)
	}
}

func TestEventWhat(t *testing.T) {
	what1 := &EventWhat{}
	what1.WithApplication(nil).
		SetMessage("application message").GetAttributes().SetString("key1", "value1").
		WithAttributes("key2", nil).
		SetString("key21", "value21")

	what1.WithLibrary("lib1", nil).
		SetMessage("lib1 message").GetAttributes().SetString("key2", "value2").
		WithAttributes("key1", nil).
		SetString("key11", "value11")

	t.Log(reflect.DeepEqual(what, what1))

	what2 := &EventWhat{}
	what2.WithApplication(what1.Application)
	what2.WithLibrary("lib1", what.Libraries.Libraries["lib1"])

	t.Log(reflect.DeepEqual(what, what2))

	t.Log(what.GetValue(&Path{Type: PathType_Application, Path: []string{"__application__", "__message__"}}))
}

func BenchmarkEventWhat(b *testing.B) {
	what1 := &EventWhat{}
	what1.WithApplication(nil).
		SetMessage("application message").GetAttributes().SetString("key1", "value1").
		WithAttributes("key2", nil).
		SetString("key21", "value21")

	what1.WithLibrary("lib1", nil).
		SetMessage("lib1 message").GetAttributes().SetString("key2", "value2").
		WithAttributes("key1", nil).
		SetString("key11", "value11")

	for i := 0; i < b.N; i++ {
		what2 := &EventWhat{}
		what2.WithApplication(what1.Application)
		what2.WithLibrary("lib1", what.Libraries.Libraries["lib1"])
	}
}
