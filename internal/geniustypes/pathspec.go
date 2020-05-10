package geniustypes

import (
	"fmt"
	"path/filepath"
	"strings"
)

const Ext = ".tmpl"

type PathSpec struct {
	In, Out string
}

func (p PathSpec) String() string { return p.In + " → " + p.Out }
func (p PathSpec) IsGoFile() bool { return filepath.Ext(p.Out) == ".go" }

func MakePathSpec(path string) (PathSpec, error) {
	p := strings.IndexByte(path, '=')
	if p == -1 {
		if filepath.Ext(path) != Ext {
			return PathSpec{}, fmt.Errorf("template file '%s' must have '%s' extension", path, Ext)
		}
		return PathSpec{path, path[:len(path)-len(Ext)]}, nil
	}

	return PathSpec{path[:p], path[p+1:]}, nil
}
