package multi_runner

import (
	"fmt"
	"testing"
	"time"
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
		}, i)
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
