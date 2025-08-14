package logger_tools

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	Key = "logger_ctx_trace"
)

type Logger struct {
	Logger     *logrus.Entry
	Formatter  Formatter
	FieldOrder []string // 记录字段添加顺序
}

// NewContext function should be updated to:
func NewContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	traceID := uuid.NewString()
	defaultFormatter := &DefaultFormatter{Split: "|"}
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(defaultFormatter)
	logger := &Logger{
		Logger:     log.WithField("trace_id", traceID),
		Formatter:  defaultFormatter,
		FieldOrder: []string{"trace_id"}, // 初始化字段顺序
	}
	return context.WithValue(ctx, Key, logger)
}

// SetCtxLoggerFormatter 设置日志格式化器
func SetCtxLoggerFormatter(ctx context.Context, formatter logrus.Formatter) {
	if ctx == nil {
		return
	}
	getLogger(ctx).Logger.Logger.SetFormatter(formatter)
}

// SetCtxLoggerLevel 设置日志级别
func SetCtxLoggerLevel(ctx context.Context, level logrus.Level) {
	if ctx == nil {
		return
	}
	getLogger(ctx).Logger.Logger.SetLevel(level)
}

// SetCtxLoggerOutput 设置日志输出
func SetCtxLoggerOutput(ctx context.Context, output io.Writer) {
	if ctx == nil {
		return
	}
	getLogger(ctx).Logger.Logger.SetOutput(output)
}

// WithField 添加字段并返回新的上下文
func WithField(ctx context.Context, key string, value any) context.Context {
	if ctx == nil {
		ctx = NewContext(nil)
	}
	logger := getLogger(ctx)

	// 创建新的字段顺序列表
	newFieldOrder := make([]string, len(logger.FieldOrder))
	copy(newFieldOrder, logger.FieldOrder)

	// 如果字段不存在，添加到顺序列表
	found := false
	for _, existingKey := range newFieldOrder {
		if existingKey == key {
			found = true
			break
		}
	}
	if !found {
		newFieldOrder = append(newFieldOrder, key)
	}

	// 创建包含字段顺序的fields
	allFields := make(logrus.Fields)
	for k, v := range logger.Logger.Data {
		allFields[k] = v
	}
	allFields[key] = value
	allFields["__field_order"] = newFieldOrder

	entry := logger.Logger.WithFields(allFields)
	newLogger := &Logger{
		Logger:     entry,
		Formatter:  logger.Formatter,
		FieldOrder: newFieldOrder,
	}
	return context.WithValue(ctx, Key, newLogger)
}

// WithFields 添加多个字段并返回新的上下文
func WithFields(ctx context.Context, fields map[string]any) context.Context {
	if ctx == nil {
		ctx = NewContext(nil)
	}
	logger := getLogger(ctx)

	// 创建新的字段顺序列表
	newFieldOrder := make([]string, len(logger.FieldOrder))
	copy(newFieldOrder, logger.FieldOrder)

	// 添加新字段到顺序列表（如果不存在）
	for key := range fields {
		found := false
		for _, existingKey := range newFieldOrder {
			if existingKey == key {
				found = true
				break
			}
		}
		if !found {
			newFieldOrder = append(newFieldOrder, key)
		}
	}

	// 创建包含字段顺序的fields
	allFields := make(logrus.Fields)
	for k, v := range logger.Logger.Data {
		allFields[k] = v
	}
	for k, v := range fields {
		allFields[k] = v
	}
	allFields["__field_order"] = newFieldOrder

	entry := logger.Logger.WithFields(allFields)
	newLogger := &Logger{
		Logger:     entry,
		Formatter:  logger.Formatter,
		FieldOrder: newFieldOrder,
	}
	return context.WithValue(ctx, Key, newLogger)
}

// Info 记录 info 级别日志
func Info(ctx context.Context, args ...any) {
	if ctx == nil {
		logrus.Info(args...)
		return
	}
	logger := getLogger(ctx)
	if logger.Formatter.GetSplit() != "" {
		strArgs := make([]string, 0, len(args))
		for _, arg := range args {
			strArgs = append(strArgs, fmt.Sprintf("%v", arg))
		}
		logger.Logger.Info(strings.Join(strArgs, logger.Formatter.GetSplit()))
	} else {
		logger.Logger.Info(args...)
	}
}

// Debug 记录 debug 级别日志
func Debug(ctx context.Context, args ...any) {
	if ctx == nil {
		logrus.Debug(args...)
		return
	}
	logger := getLogger(ctx)
	if logger.Formatter.GetSplit() != "" {
		strArgs := make([]string, 0, len(args))
		for _, arg := range args {
			strArgs = append(strArgs, fmt.Sprintf("%v", arg))
		}
		logger.Logger.Debug(strings.Join(strArgs, logger.Formatter.GetSplit()))
	} else {
		logger.Logger.Debug(args...)
	}
}

// Warn 记录 warn 级别日志
func Warn(ctx context.Context, args ...any) {
	if ctx == nil {
		logrus.Warn(args...)
		return
	}
	logger := getLogger(ctx)
	if logger.Formatter.GetSplit() != "" {
		strArgs := make([]string, 0, len(args))
		for _, arg := range args {
			strArgs = append(strArgs, fmt.Sprintf("%v", arg))
		}
		logger.Logger.Warn(strings.Join(strArgs, logger.Formatter.GetSplit()))
	} else {
		logger.Logger.Warn(args...)
	}
}

// Error 记录 error 级别日志
func Error(ctx context.Context, args ...any) {
	if ctx == nil {
		logrus.Error(args...)
		return
	}
	logger := getLogger(ctx)
	if logger.Formatter.GetSplit() != "" {
		strArgs := make([]string, 0, len(args))
		for _, arg := range args {
			strArgs = append(strArgs, fmt.Sprintf("%v", arg))
		}
		logger.Logger.Error(strings.Join(strArgs, logger.Formatter.GetSplit()))
	} else {
		logger.Logger.Error(args...)
	}
}

// Fatal 记录 fatal 级别日志
func Fatal(ctx context.Context, args ...any) {
	if ctx == nil {
		logrus.Fatal(args...)
		return
	}
	logger := getLogger(ctx)
	if logger.Formatter.GetSplit() != "" {
		strArgs := make([]string, 0, len(args))
		for _, arg := range args {
			strArgs = append(strArgs, fmt.Sprintf("%v", arg))
		}
		logger.Logger.Fatal(strings.Join(strArgs, logger.Formatter.GetSplit()))
	} else {
		logger.Logger.Fatal(args...)
	}
}

// Panic 记录 panic 级别日志
func Panic(ctx context.Context, args ...any) {
	if ctx == nil {
		logrus.Panic(args...)
		return
	}
	logger := getLogger(ctx)
	if logger.Formatter.GetSplit() != "" {
		strArgs := make([]string, 0, len(args))
		for _, arg := range args {
			strArgs = append(strArgs, fmt.Sprintf("%v", arg))
		}
		logger.Logger.Panic(strings.Join(strArgs, logger.Formatter.GetSplit()))
	} else {
		logger.Logger.Panic(args...)
	}
}

// getLogger 从 context 获取 logger
func getLogger(ctx context.Context) *Logger {
	if ctx == nil {
		return &Logger{
			Logger:     logrus.NewEntry(logrus.StandardLogger()),
			Formatter:  &DefaultFormatter{Split: "|"},
			FieldOrder: []string{},
		}
	}

	if logger, ok := ctx.Value(Key).(*Logger); ok {
		return logger
	}
	return &Logger{
		Logger:     logrus.NewEntry(logrus.StandardLogger()),
		Formatter:  &DefaultFormatter{Split: "|"},
		FieldOrder: []string{},
	}
}
