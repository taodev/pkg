package types

import (
	"encoding/hex"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestBinary(t *testing.T) {
	const (
		base64Data = "SGVsbG8gQmluYXJ5IQ"
		hexData    = "48656c6c6f2042696e61727921"
		strData    = "Hello Binary!"
	)
	var b Binary
	err := b.UnmarshalYAML(&yaml.Node{
		Tag:   "!!binary",
		Value: base64Data,
	})
	assert.NoError(t, err)
	assert.Equal(t, []byte(strData), []byte(b))
	assert.Equal(t, hexData, hex.EncodeToString(b))
	assert.Equal(t, base64Data, b.String())

	err = b.UnmarshalYAML(&yaml.Node{
		Tag:   "!!hex",
		Value: hexData,
	})
	assert.NoError(t, err)
	assert.Equal(t, []byte(strData), []byte(b))

	// unknown tag
	assert.Error(t, b.UnmarshalYAML(&yaml.Node{
		Tag:   "!!int",
		Value: base64Data,
	}))

	// error value
	assert.Error(t, b.UnmarshalYAML(&yaml.Node{
		Tag:   "!!binary",
		Value: "error value",
	}))

	yamlStruct := &struct {
		Value1 Binary `yaml:"value1"`
		Value2 Binary `yaml:"value2"`
		Value3 Binary `yaml:"value3"`
		Value4 Binary `yaml:"value4"`
		Value5 Binary `yaml:"value5"`
	}{}
	data, err := os.ReadFile("data/binary.yaml")
	require.NoError(t, err)
	err = yaml.Unmarshal(data, yamlStruct)
	require.NoError(t, err)
	assert.Equal(t, []byte(strData), yamlStruct.Value1.Bytes())
	assert.Equal(t, []byte(strData), yamlStruct.Value2.Bytes())
	assert.Equal(t, []byte(strData), yamlStruct.Value3.Bytes())
	assert.Equal(t, []byte(strData), yamlStruct.Value4.Bytes())
	assert.Equal(t, []byte(strData), yamlStruct.Value5.Bytes())

	data, err = yaml.Marshal(yamlStruct.Value1)
	require.NoError(t, err)
	log.Printf("%s", data)
	assert.Contains(t, string(data), base64Data)

	var b1 Binary
	err = b1.Parse(base64Data)
	require.NoError(t, err)
	assert.Equal(t, []byte(strData), b1.Bytes())

	err = b1.Parse("error data")
	require.Error(t, err)
}
