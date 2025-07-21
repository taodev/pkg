package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYAML 函数用于从指定文件中读取 YAML 数据并将其反序列化到给定的变量中。
// 参数 filename 是要读取的 YAML 文件的路径。
// 参数 v 是用于存储反序列化结果的变量指针，类型为 any。
// 返回值为错误信息，如果读取或反序列化过程中出现错误，则返回相应的错误；否则返回 nil。
func LoadYAML(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

// SaveYAML 函数用于将给定的变量序列化为 YAML 格式并保存到指定文件中。
// 参数 filename 是要保存的 YAML 文件的路径。
// 参数 v 是要序列化的变量，类型为 any。
// 返回值为错误信息，如果序列化或写入文件过程中出现错误，则返回相应的错误；否则返回 nil。
func SaveYAML(filename string, v any) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
