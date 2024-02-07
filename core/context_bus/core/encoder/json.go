package encoder

type jsonEncoder struct{}

var JSONEncoder jsonEncoder

func (e *jsonEncoder) BeginObject(buf []byte) []byte {
	return append(buf, '{')
}

func (e *jsonEncoder) EndObject(buf []byte) []byte {
	return append(buf, '}')
}

func (e *jsonEncoder) BeginString(buf []byte) []byte {
	return append(buf, '"')
}

func (e *jsonEncoder) EndString(buf []byte) []byte {
	return append(buf, '"')
}

func (e *jsonEncoder) AppendKey(dst []byte, key string) []byte {
	if dst[len(dst)-1] != '{' {
		dst = append(dst, ',')
	}

	return append(e.AppendString(dst, key), ':')
}

func (e *jsonEncoder) AppendString(dst []byte, str string) []byte {
	dst = append(dst, '"')
	// todo: check escaped
	dst = append(dst, str...)

	return append(dst, '"')
}
