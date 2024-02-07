package code_generator

import (
	"testing"
)

func TestGetCodeInfoBasic(t *testing.T) {
	info := CodeInfoStorage.GetCodeInfoBasic(CodeInfo_Test, CodeInfoSkip)
	t.Logf("%+v", info)
}

func BenchmarkGetCodeInfoBasic_Generated(b *testing.B) {
	fc := func() {
		info := CodeInfoStorage.GetCodeInfoBasic(CodeInfo_Test, CodeInfoSkip + 1)
		if info == nil {
			b.Error(info)
		}
	}

	for i := 0; i < b.N; i++ {
		fc()
	}
}

func BenchmarkGetCodeInfoBasic(b *testing.B) {
	fc := func() {
		info := getCodeInfoBasic(2)
		if info == nil {
			b.Error(info)
		}
	}

	for i := 0; i < b.N; i++ {
		fc()
	}
}
