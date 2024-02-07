package code_generator

import (
	"github.com/AleckDarcy/reload/core/context_bus/helper"
	"regexp"
	"strings"
)

// User message contains placeholders as the keys to attributes.
// e.g., "this message contains information from lib1: {lib1.key1} and lib2: {lib2.message}"
// parser converts this message to "this message contains information from lib1: %s and lib2: %s",
// converts placeholders to paths (e.g., lib1->key1),
// and fetches the values from EventWhat (e.g., EventWhat.GetValue(Path)
// todo: code generator. Can be replaced by one-time effort.

type PathType uint64

const (
	_ PathType = iota
	PathApplication
	PathLibrary
)

type Path struct {
	Type PathType
	Path []string
}

var PathHelper = &Path{}
var pathReg = regexp.MustCompile(`\${[^}]*}`)

func (g *Path) ParseMessage(msg string) (string, []*Path) {
	idxs := pathReg.FindAllStringIndex(msg, -1)
	fmtLen := len(msg)
	for _, idx := range idxs {
		fmtLen += 2 - (idx[1] - idx[0])  // 2 bytes for %s
	}

	fmtBytes := make([]byte, fmtLen)
	paths := make([]*Path, 0, len(idxs))
	id, offset := 0, 0
	for _, idx := range idxs {

		if idx[0] != id {
			copy(fmtBytes[offset:], msg[id: idx[0]])
			offset += idx[0] - id
		}

		copy(fmtBytes[offset:], "%s")
		paths = append(paths, PathHelper.ParsePath(msg[idx[0]: idx[1]]))

		offset += 2 // 2 bytes for %s
		id = idx[1]
	}

	return helper.BytesToString(fmtBytes), paths
}

func (g *Path) ParsePath(str string) *Path {
	path := strings.Split(str, ".")

	if path[0] == "_" { // application. e.g., _.key
		return &Path{Type: PathApplication, Path: path[1:]}
	}

	return &Path{Type: PathLibrary, Path: path}
}
