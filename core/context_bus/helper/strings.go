package helper

import (
	"reflect"
	"unsafe"
)

func StringToBytes(str string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))

	bytesHeader := &reflect.SliceHeader{
		Data: stringHeader.Data,
		Len: stringHeader.Len,
		Cap: stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}

func StringToBytesV2(str string) (bytes []byte) {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))

	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	bytesHeader.Data = stringHeader.Data
	bytesHeader.Len = stringHeader.Len
	bytesHeader.Cap = stringHeader.Len

	return
}

func BytesToString(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))

	stringHeader := &reflect.StringHeader{
		Data: bytesHeader.Data,
		Len: bytesHeader.Len,
	}

	return *(*string)(unsafe.Pointer(stringHeader))
}