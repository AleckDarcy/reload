package proto

import (
	"reflect"
	"testing"
)

var attrs1 = &Attributes{
	Attrs: map[string]*AttributeValue{
		"key1": {
			Type:  AttributeValueType_AttributeValueStr,
			Value: &AttributeValue_Str{Str: "value1"},
		},
		"key2": {
			Type: AttributeValueType_AttributeValueAttr,
			Value: &AttributeValue_Struct{
				Struct: &Attributes{
					Attrs: map[string]*AttributeValue{
						"key21": {
							Type:  AttributeValueType_AttributeValueStr,
							Value: &AttributeValue_Str{Str: "value21"},
						},
					},
				},
			},
		},
	},
}

var attrs2 = &Attributes{
	Attrs: map[string]*AttributeValue{
		"key1": {
			Type:  AttributeValueType_AttributeValueStr,
			Value: &AttributeValue_Str{Str: "value2"},
		},
		"key3": {
			Type: AttributeValueType_AttributeValueAttr,
			Value: &AttributeValue_Struct{
				Struct: &Attributes{
					Attrs: map[string]*AttributeValue{
						"key31": {
							Type:  AttributeValueType_AttributeValueStr,
							Value: &AttributeValue_Str{Str: "value31"},
						},
					},
				},
			},
		},
	},
}

var attrs3 = &Attributes{
	Attrs: map[string]*AttributeValue{
		"key1": {
			Type:  AttributeValueType_AttributeValueStr,
			Value: &AttributeValue_Str{Str: "value1"},
		},
		"key2": {
			Type: AttributeValueType_AttributeValueAttr,
			Value: &AttributeValue_Struct{
				Struct: &Attributes{
					Attrs: map[string]*AttributeValue{
						"key21": {
							Type:  AttributeValueType_AttributeValueStr,
							Value: &AttributeValue_Str{Str: "value21"},
						},
					},
				},
			},
		},
		"key3": {
			Type: AttributeValueType_AttributeValueAttr,
			Value: &AttributeValue_Struct{
				Struct: &Attributes{
					Attrs: map[string]*AttributeValue{
						"key31": {
							Type:  AttributeValueType_AttributeValueStr,
							Value: &AttributeValue_Str{Str: "value31"},
						},
					},
				},
			},
		},
	},
}

func TestAttributes(t *testing.T) {
	attrs1.Merge(attrs2)

	if !reflect.DeepEqual(attrs1, attrs3) {
		t.Error("fail")
	}
}

func TestAttributes_GetValue(t *testing.T) {
	t.Log(attrs3.GetValue([]string{"key2", "key21"}))
}

func BenchmarkAttributes_GetValue(b *testing.B) {
	key2 := attrs3.Attrs["key2"]
	key3 := key2.Value.(*AttributeValue_Struct).Struct
	for i := 0; i < b.N; i++ {
		if key2.Type == AttributeValueType_AttributeValueAttr {
			key3 = key2.Value.(*AttributeValue_Struct).Struct
		}
		//attrs3.GetValue([]string{"key2", "key21"})
	}

	b.Log(key3)
}


