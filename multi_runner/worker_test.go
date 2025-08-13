package multi_runner

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/otkinlife/go_tools/logger_tools"
	"github.com/sirupsen/logrus"
)

func TestRun(t *testing.T) {
	r := NewRunner(10)
	for i := 0; i < 100; i++ {
		err := r.AddJob(func(data any) JobRet {
			time.Sleep(10 * time.Second)
			return JobRet{
				Err:  nil,
				Data: fmt.Sprintf("hello %v", data),
			}
		}, i, 1)
		if err != nil {
			t.Fatal(err)
			return
		}
	}

	r.Run()
	r.HandleResultsWithStream(func(ret JobRet) {
		t.Log(ret)
	})
	return
}

func TestRunWithCtx(t *testing.T) {
	ctx := logger_tools.NewContext(context.Background())
	ctx = logger_tools.WithField(ctx, "test", "multi_runner")
	logger_tools.SetCtxLoggerLevel(ctx, logrus.DebugLevel)
	logger_tools.Info(ctx, "Starting test for multi_runner with context")
	r := NewRunnerWithCtx(ctx, 10)
	for i := 0; i < 100; i++ {
		err := r.AddJob(func(ctx context.Context, data any) JobRet {
			time.Sleep(10 * time.Second)
			logger_tools.Info(ctx, "Running job with data", data)
			return JobRet{
				Err:  nil,
				Data: fmt.Sprintf("hello %v", data),
			}
		}, i, 1)
		if err != nil {
			t.Fatal(err)
			return
		}
	}

	r.Run()
	r.HandleResultsWithStream(ctx, func(ctx context.Context, ret JobRet) {
		t.Log(ret)
	})
	return
}
