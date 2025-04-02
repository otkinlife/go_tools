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
	if len(entry.Data) > 0 {
		for k, v := range entry.Data {
			userMessage = append(userMessage, fmt.Sprintf("%s=%v", k, v))
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
