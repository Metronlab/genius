package tmpl

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"text/template"

	"github.com/Metronlab/genius/internal/geniusio"
	"github.com/Metronlab/genius/internal/geniustypes"
	jsoniter "github.com/json-iterator/go"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Tmpl(dataPath string, values geniustypes.ValuesMap, args []string,
	goImportsEnable bool, write geniusio.GenerationWriteFunc) error {
	var err error
	entries := specEnvironment{
		Values: values,
	}

	if goImportsEnable {
		if _, err := exec.LookPath("goimports"); err != nil {
			return fmt.Errorf("failed to find goimports: %w", err)
		}
		formatter = formatSource
	} else {
		formatter = format.Source
	}

	if len(args) == 0 {
		return errors.New("no tmpl files specified")
	}

	entries.Data, err = geniusio.ReadFileData(dataPath)
	if err != nil {
		return fmt.Errorf("impossible to read input data file: %w", err)
	}

	specs := make([]geniustypes.TmplSpecPaths, len(args))
	for i, p := range args {
		if specs[i], err = geniustypes.MakePathSpec(p); err != nil {
			return err
		}
	}

	for _, spec := range specs {
		log.Printf("generating %s\n", spec)
		generated, err := generate(entries, spec)
		if err != nil {
			return fmt.Errorf("error execution spec %s: %w", spec.In, err)
		}

		if err := write(spec, generated); err != nil {
			return err
		}
	}

	return nil
}

type specEnvironment struct {
	Data   interface{}
	Values geniustypes.ValuesMap
}

var funcs = template.FuncMap{
	"stringsLower": strings.ToLower,
	"stringsUpper": strings.ToUpper,
	"stringsTitle": strings.Title,
	"marshalJson": func(v interface{}) string {
		res, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		return string(res)
	},
	"marshalYaml": func(v interface{}) string {
		res, err := yaml.Marshal(v)
		if err != nil {
			panic(err)
		}
		return string(res)
	},
	"marshalToml": func(v interface{}) string {
		res, err := toml.Marshal(v)
		if err != nil {
			panic(err)
		}
		return string(res)
	},
}

func generate(data interface{}, spec geniustypes.TmplSpecPaths) ([]byte, error) {
	var (
		t   *template.Template
		err error
	)

	tmplContent, err := ioutil.ReadFile(spec.In)
	if err != nil {
		return nil, fmt.Errorf("impossible to read input template data: %w", err)
	}

	t, err = template.New("gen").Funcs(funcs).Parse(string(tmplContent))
	if err != nil {
		return nil, fmt.Errorf("error processing template '%s': %w", spec.In, err)
	}

	var buf bytes.Buffer
	if spec.IsGoFile() {
		// preamble
		if _, err := fmt.Fprintf(&buf, "// Code generated by %s. DO NOT EDIT.\n", spec.In); err != nil {
			return nil, fmt.Errorf("impossible to write header: %w", err)
		}
		if _, err := fmt.Fprintln(&buf); err != nil {
			return nil, fmt.Errorf("impossible to write header: %w", err)
		}
	}
	err = t.Execute(&buf, data)
	if err != nil {
		return nil, fmt.Errorf("error executing template '%s': %w", spec.In, err)
	}

	generated := buf.Bytes()
	if spec.IsGoFile() {
		generated, err = formatter(generated)
		if err != nil {
			return nil, fmt.Errorf("error formatting template '%s': %w", spec.In, err)
		}
	}

	return generated, nil
}

var (
	formatter func([]byte) ([]byte, error)
)

func formatSource(in []byte) ([]byte, error) {
	r := bytes.NewReader(in)
	cmd := exec.Command("goimports")
	cmd.Stdin = r
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("error running goimports: %s", string(ee.Stderr))
		}
		return nil, fmt.Errorf("error running goimports: %s", string(out))
	}

	return out, nil
}
