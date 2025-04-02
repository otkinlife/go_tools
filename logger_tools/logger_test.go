package logger_tools

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	ctx := NewContext(context.Background())
	SetCtxLoggerLevel(ctx, logrus.DebugLevel)
	Info(ctx, "Test message", "info")
	Debug(ctx, "Test message", "debug")
	Warn(ctx, "Test message", "warn")
}
