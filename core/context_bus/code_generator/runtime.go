package code_generator

import (
	"runtime"
	"sync"
)

// CodeInfo Basic code base information containing
// todo: code generator. Can be replaced by one-time effort.
// new(sync.Once).do(func(){ info = GetCodeInfo })

type codeInfoStorage struct {
	infos []CodeInfo
}

// CodeInfoID is a unique identifiers of places where implementations of observation are re-written by code generator.
type CodeInfoID int

const (
	CodeInfo_ CodeInfoID = iota
	CodeInfo_Test

	CodeInfo_End
)

const CodeInfoSkip = 5

var CodeInfoStorage codeInfoStorage

type CodeInfoBasic struct {
	pc   uintptr
	name string
	file string
	line int
}

type CodeInfo struct {
	sync.Once
	basic CodeInfoBasic
}

// getCodeInfoBasic gets static code information such as file and line where (re-written) observation happens.
// todo: generator can provide code information for original code through AST.
func getCodeInfoBasic(skip int) *CodeInfoBasic {
	pc, file, line, _ := runtime.Caller(skip)

	return &CodeInfoBasic{
		pc:   pc,
		name: runtime.FuncForPC(pc).Name(),
		file: file,
		line: line,
	}
}

// GetCodeInfoBasic indexed by CodeInfoID (see description)
func (f *codeInfoStorage) GetCodeInfoBasic(id CodeInfoID, skip int) *CodeInfoBasic {
	if id <= 0 || id >= CodeInfo_End {
		panic("unreachable code")
	}

	info := &f.infos[id]
	info.Once.Do(func() {
		info.basic = *getCodeInfoBasic(skip)
	})

	return &info.basic
}
