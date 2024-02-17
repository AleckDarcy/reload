package proto

// global variables for testing only

var Test_EventMessage_Rest = &EventMessage{
	Attrs: &Attributes{
		Attrs: map[string]*AttributeValue{
			"from": {
				Type: AttributeValueType_AttributeValueStr,
				Str:  "SenderA",
			},
			"key": {
				Type: AttributeValueType_AttributeValueStr,
				Str:  "This a string attribute",
			},
			"key_": {
				Type: AttributeValueType_AttributeValueStr,
				Str:  "This another string attribute",
			},
		},
	},
}
