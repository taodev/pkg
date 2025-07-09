package types

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	// BinaryTag 二进制编码标签
	BinaryTag = "!!binary"
	// BinaryStrTag 二进制字符串编码标签
	BinaryStrTag = "!!str"
	// BinaryBase64Tag 二进制Base64编码标签
	BinaryBase64Tag = "!!base64"
	// BinaryHexTag 二进制Base64原始编码标签
	BinaryHexTag = "!!hex"
	BinarySeqTag = "!!seq"
)

type Binary []byte

// Yaml 编码
func (b Binary) MarshalYAML() (any, error) {
	return yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: base64.StdEncoding.EncodeToString(b),
	}, nil
}

func (b *Binary) UnmarshalYAML(value *yaml.Node) error {
	var decoded []byte
	var err error
	switch value.Tag {
	case BinaryTag, BinaryStrTag, BinaryBase64Tag:
		decoded, err = base64.StdEncoding.DecodeString(value.Value)
	case BinaryHexTag:
		decoded, err = hex.DecodeString(value.Value)
	case BinarySeqTag:
		err = value.Decode(&decoded)
	default:
		return fmt.Errorf("unexpected tag: %s", value.Tag)
	}
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
