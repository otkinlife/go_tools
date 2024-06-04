package multi_runner

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

const (
	StatusWait = 0
	StatusRun  = 1
	StatusEnd  = 2
)

type JobExecute func(data any) JobRet
type JobRetHandler func(ret JobRet)

type JobRet struct {
	Err  error
	Data any
}

// Job 任务
type Job struct {
	ID        string     // 任务ID
	Execute   JobExecute // 执行函数
	RunParams any        // 执行数据
	RunStatus int        // 运行状态: 0表示排队中，1表示运行中，2表示已结束
	RunRets   JobRet     // 执行结果
}

type Runner struct {
	maxSize     int //最大同时并发量
	jobs        sync.Map
	jobsCount   int
	sem         chan int
	wg          sync.WaitGroup
	isRunnerEnd chan int
	isHandled   bool
	results     chan JobRet
}

func NewRunner(maxSize int) *Runner {
	return &Runner{
		wg:          sync.WaitGroup{},
		maxSize:     maxSize,
		jobs:        sync.Map{},
		sem:         make(chan int, maxSize),
		results:     make(chan JobRet),
		isRunnerEnd: make(chan int, 1),
	}
}

func (r *Runner) AddJob(handler JobExecute, runParams any) error {
	jobID := uuid.NewString()
	r.jobs.Store(jobID, &Job{
		ID:        jobID,
		Execute:   handler,
		RunStatus: StatusWait,
		RunParams: runParams,
	})
	r.jobsCount++
	return nil
}

func (r *Runner) Run() {
	r.results = make(chan JobRet, r.jobsCount)
	if r.jobsCount < r.maxSize {
		r.maxSize = r.jobsCount
		r.sem = make(chan int, r.maxSize)
	}
	r.wg.Add(r.jobsCount)
	r.jobs.Range(func(key, value any) bool {
		job := value.(*Job)
		job.RunStatus = StatusRun
		go func(id string, handler JobExecute, params any) {
			ret := JobRet{
				Err:  nil,
				Data: nil,
			}
			r.sem <- StatusRun
			defer func() {
				if err := recover(); err != nil {
					ret.Err = fmt.Errorf("panic: %v", err)
				}
				r.results <- ret
				<-r.sem
				job.RunStatus = StatusEnd
				r.wg.Done()
			}()
			ret = handler(params)
		}(job.ID, job.Execute, job.RunParams)
		return true
	})
}

func (r *Runner) HandleResultsWithStream(handler JobRetHandler) {
	if r.isHandled {
		return
	}
	r.isHandled = true
	go func() {
		r.wg.Wait()
		close(r.results)
		r.isRunnerEnd <- 1
	}()
	// 实时监听结果，直到所有任务完成
	for {
		select {
		case ret, ok := <-r.results:
			if !ok {
				return
			}
			handler(ret)
		case isRunnerEnd := <-r.isRunnerEnd:
			if isRunnerEnd == 1 {
				return
			}
		}
	}
}

func (r *Runner) HandleAllResultsWith(handler JobRetHandler) {
	if r.isHandled {
		return
	}
	r.isHandled = true
	r.wg.Wait()
	close(r.results)
	for ret := range r.results {
		handler(ret)
	}
}
