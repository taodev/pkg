package types

import (
	"encoding/base64"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Binary []byte

// Yaml 编码
func (b Binary) MarshalYAML() (any, error) {
	return yaml.Node{
		Kind:  yaml.ScalarNode,
		Tag:   "!!binary",
		Value: base64.StdEncoding.EncodeToString(b),
	}, nil
}

func (b *Binary) UnmarshalYAML(value *yaml.Node) error {
	if value.Tag != "!!binary" && value.Tag != "!!str" {
		return fmt.Errorf("unexpected tag: %s", value.Tag)
	}
	decoded, err := base64.StdEncoding.DecodeString(value.Value)
	if err != nil {
		return err
	}
	*b = decoded
	return nil
}

func (b Binary) String() string {
	return base64.StdEncoding.EncodeToString(b)
}

func (b Binary) Bytes() []byte {
	return b
}
