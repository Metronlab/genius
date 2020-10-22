package geniustypes

import (
	"fmt"
	"path/filepath"
	"strings"
)

const Ext = ".tmpl"

type TmplSpecPaths struct {
	In, Out string
}

func (p TmplSpecPaths) String() string { return p.In + " → " + p.Out }
func (p TmplSpecPaths) IsGoFile() bool { return filepath.Ext(p.Out) == ".go" }

func MakePathSpec(outputPrefix, path string) (TmplSpecPaths, error) {
	p := strings.IndexByte(path, '=')
	if p == -1 {
		if filepath.Ext(path) != Ext {
			return TmplSpecPaths{}, fmt.Errorf("template file '%s' must have '%s' extension", path, Ext)
		}
		return TmplSpecPaths{path, path[:len(path)-len(Ext)]}, nil
	}

	return TmplSpecPaths{path[:p], outputPrefix + path[p+1:]}, nil
}
