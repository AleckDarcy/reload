package proto

// global variables for testing only

var Test_EventMessage_Rest = &EventMessage{
	Attrs: &Attributes{
		Attrs: map[string]*AttributeValue{
			"from": {
				Type: AttributeValueType_AttributeValueStr,
				Str:  "SenderA",
			},
			"method": {
				Type: AttributeValueType_AttributeValueStr,
				Str:  "POST",
			},
			"handler": {
				Type: AttributeValueType_AttributeValueStr,
				Str:  "/request_b",
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
