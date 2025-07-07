package defaults

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYAML(filename string, v any) error {
	if err := Set(v); err != nil {
		return err
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

func LoadJSON(filename string, v any) error {
	if err := Set(v); err != nil {
		return err
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
