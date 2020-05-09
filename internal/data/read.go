package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var ErrUnsupportedEncoding = errors.New("unsupported encoding")

//ReadFileData decode toml, yaml or json based on file extension
func ReadFileData(path string) (interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var v interface{}

	switch ext := filepath.Ext(path); ext {
	case ".json", ".JSON":
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("invalid JSON data: %w", err)
		}
	case ".yaml", ".YAML", ".yml", ".YML":
		if err := yaml.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("invalid YAML data: %w", err)
		}
	case ".toml", ".TOML":
		if err := toml.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("invalid TOML data: %w", err)
		}
	default:
		return nil, fmt.Errorf("decoding file with extension %s: %w", ext, ErrUnsupportedEncoding)
	}

	return v, nil
}
