package geniusio

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

var ErrUnsupportedEncoding = errors.New("unsupported encoding")

//ReadFileData decode toml, yaml or json based on file extension
func ReadFileData(path, format string) (interface{}, error) {
	var unmarshal func(data []byte, v interface{}) error
	var err error
	if path == "" {
		if format == "" {
			format = "txt"
		}
		unmarshal, err = makeUnmarshaller(fmt.Sprintf(".%s", format))
	} else {
		unmarshal, err = makeUnmarshaller(filepath.Ext(path))
	}
	if err != nil {
		return nil, err
	}

	if path == "" {
		return readStdIn(unmarshal)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var v interface{}
	if err := unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("invalid data: %w", err)
	}
	return v, nil
}

func readStdIn(unmarshall func(data []byte, v interface{}) error) (v interface{}, err error) {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadBytes(0)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	err = unmarshall(data, &v)
	return
}

func makeUnmarshaller(extension string) (func(data []byte, v interface{}) error, error) {
	switch extension {
	case ".json", ".JSON":
		return json.Unmarshal, nil
	case ".yaml", ".YAML", ".yml", ".YML":
		return yaml.Unmarshal, nil
	case ".toml", ".TOML":
		return toml.Unmarshal, nil
	case ".txt", ".text":
		return func(data []byte, v interface{}) error {
			rv := reflect.ValueOf(v)
			if rv.Kind() != reflect.Ptr || rv.IsNil() {
				panic("only accept interface{}")
			}
			rv.Elem().Set(reflect.ValueOf(string(data)))
			return nil
		}, nil
	default:
		return nil, fmt.Errorf("decoding file with format \"%s\": %w", extension, ErrUnsupportedEncoding)
	}
}
