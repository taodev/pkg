package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYAML(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

func SaveYAML(filename string, v any) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
