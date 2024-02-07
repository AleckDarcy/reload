package helper

import (
	"testing"
)

func BenchmarkStringToBytes(b *testing.B) {
	var str = "testtesttest"
	for i := 0; i < b.N; i++ {
		StringToBytes(str)
	}
}

func BenchmarkStringToBytesV2(b *testing.B) {
	var str = "testtesttest"
	for i := 0; i < b.N; i++ {
		StringToBytesV2(str)
	}
}

func BenchmarkCopy(b *testing.B) {
	str := BytesToString(make([]byte, 100))
	bytes := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		copy(bytes, str[90: 100])
	}
}
