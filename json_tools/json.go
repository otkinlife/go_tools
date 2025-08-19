package json_tools

import (
	"encoding/json"
	"fmt"

	"github.com/otkinlife/go_tools/logger_tools"
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

// UnmarshalJsonWithoutError 尝试将JSON字符串解码到提供的接口中，但不返回错误
// 如果解码失败，将打印错误信息
// 注意：这种方式可能会导致数据不完整或错误，因此仅在你确定数据
func UnmarshalJsonWithoutError(data string, v any) {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		logger_tools.Error(nil, "解码 JSON 失败", err, data)
	}
}

// MarshalJsonWithoutError 尝试将提供的接口编码为JSON字符串，但不返回错误
func MarshalJsonWithoutError(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		logger_tools.Error(nil, "编码 JSON 失败", err, v)
		return ""
	}
	return string(data)
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
