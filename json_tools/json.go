package json_tools

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJson 将JSON字符串解码到提供的接口中
func UnmarshalJson(data string, v any) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

// MarshalJson 将提供的接口编码为JSON字符串
func MarshalJson(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(data), nil
}

// MarshalJsonPretty 将提供的接口编码为格式化的JSON字符串
func MarshalJsonPretty(v any) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON with indent: %w", err)
	}
	return string(data), nil
}

// IsValidJson 检查提供的字符串是否为有效的JSON格式
func IsValidJson(data string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(data), &js) == nil
}
