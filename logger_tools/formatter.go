package logger_tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Formatter interface {
	logrus.Formatter
	GetSplit() string
	SetSplit(split string)
}

type DefaultFormatter struct {
	Split string
}

func (f *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	now := time.Now()
	userMessage := []string{
		now.Format(time.DateTime),
		fmt.Sprintf("%d", now.UnixMilli()),
		fmt.Sprintf("level=%s", entry.Level.String()),
	}

	// 在level后面添加trace_id（如果存在）
	if traceID, exists := entry.Data["trace_id"]; exists {
		userMessage = append(userMessage, fmt.Sprintf("trace_id=%v", traceID))
	}

	// 尝试从entry.Data获取字段顺序
	if fieldOrder, ok := entry.Data["__field_order"]; ok {
		if order, ok := fieldOrder.([]string); ok {
			// 按照字段添加顺序输出（跳过trace_id，已经在前面添加了）
			for _, key := range order {
				if key != "__field_order" && key != "trace_id" {
					if v, exists := entry.Data[key]; exists {
						userMessage = append(userMessage, fmt.Sprintf("%s=%v", key, v))
					}
				}
			}
		} else {
			// 回退到原始方式
			for k, v := range entry.Data {
				if k != "__field_order" && k != "trace_id" {
					userMessage = append(userMessage, fmt.Sprintf("%s=%v", k, v))
				}
			}
		}
	} else if len(entry.Data) > 0 {
		// 没有顺序信息，使用原始方式（跳过trace_id）
		for k, v := range entry.Data {
			if k != "trace_id" {
				userMessage = append(userMessage, fmt.Sprintf("%s=%v", k, v))
			}
		}
	}

	if entry.Message != "" {
		userMessage = append(userMessage, entry.Message)
	}
	msg := strings.Join(userMessage, f.Split) + "\n"
	return []byte(msg), nil
}

func (f *DefaultFormatter) GetSplit() string {
	return f.Split
}

func (f *DefaultFormatter) SetSplit(split string) {
	f.Split = split
}
