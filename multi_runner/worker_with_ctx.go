package multi_runner

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type JobExecuteWithCtx func(ctx context.Context, data any) JobRet
type JobRetHandlerWithCtx func(ctx context.Context, ret JobRet)
type JobRetOutputHandlerWithCtx func(ctx context.Context, ret JobRet, output any)

// JobWithCtx 支持上下文的任务
type JobWithCtx struct {
	ID        string            // 任务ID
	Execute   JobExecuteWithCtx // 执行函数（支持context）
	RunParams any               // 执行数据
	RunStatus int               // 运行状态: 0表示排队中，1表示运行中，2表示已结束
	RunRets   JobRet            // 执行结果
	Retry     int               // 重试次数
	MaxRetry  int               // 最大重试次数
}

type RunnerWithCtx struct {
	ctx         context.Context
	cancel      context.CancelFunc
	maxSize     int //最大同时并发量
	jobs        sync.Map
	jobsCount   int
	sem         chan int
	wg          sync.WaitGroup
	isRunnerEnd chan int
	isHandled   bool
	results     chan JobRet
	mu          sync.RWMutex // 保护 isHandled 字段
}

// NewRunnerWithCtx 创建支持上下文的新Runner
func NewRunnerWithCtx(ctx context.Context, maxSize int) *RunnerWithCtx {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	return &RunnerWithCtx{
		ctx:         ctxWithCancel,
		cancel:      cancel,
		wg:          sync.WaitGroup{},
		maxSize:     maxSize,
		jobs:        sync.Map{},
		sem:         make(chan int, maxSize),
		results:     make(chan JobRet),
		isRunnerEnd: make(chan int, 1),
	}
}

// AddJob 添加支持上下文的任务
func (r *RunnerWithCtx) AddJob(handler JobExecuteWithCtx, runParams any, maxRetry int) error {
	if maxRetry <= 0 {
		maxRetry = 1
	}
	jobID := uuid.NewString()
	r.jobs.Store(jobID, &JobWithCtx{
		ID:        jobID,
		Execute:   handler,
		RunStatus: StatusWait,
		RunParams: runParams,
		MaxRetry:  maxRetry,
		Retry:     0,
	})
	r.jobsCount++
	return nil
}

// Run 运行所有任务
func (r *RunnerWithCtx) Run() {
	r.results = make(chan JobRet, r.jobsCount)
	if r.jobsCount < r.maxSize {
		r.maxSize = r.jobsCount
		r.sem = make(chan int, r.maxSize)
	}
	r.wg.Add(r.jobsCount)

	r.jobs.Range(func(key, value any) bool {
		job := value.(*JobWithCtx)
		job.RunStatus = StatusRun
		go func(job *JobWithCtx) {
			ret := JobRet{
				Err:  nil,
				Data: nil,
			}

			// 检查上下文是否已经被取消
			if r.ctx.Err() != nil {
				ret.Err = r.ctx.Err()
				r.results <- ret
				job.RunStatus = StatusEnd
				r.wg.Done()
				return
			}

			select {
			case r.sem <- StatusRun:
				// 成功获取信号量
			case <-r.ctx.Done():
				// 上下文被取消
				ret.Err = r.ctx.Err()
				r.results <- ret
				job.RunStatus = StatusEnd
				r.wg.Done()
				return
			}

			defer func() {
				if err := recover(); err != nil {
					ret.Err = fmt.Errorf("panic: %v", err)
				}
				r.results <- ret
				<-r.sem
				job.RunStatus = StatusEnd
				r.wg.Done()
			}()

			// 执行任务重试逻辑
			for job.Retry < job.MaxRetry {
				// 检查上下文是否被取消
				if r.ctx.Err() != nil {
					ret.Err = r.ctx.Err()
					break
				}

				job.Retry++
				ret = job.Execute(r.ctx, job.RunParams)
				if ret.Err == nil {
					break
				}
			}
		}(job)
		return true
	})

	go func() {
		r.wg.Wait()
		close(r.results)
	}()
}

// HandleResultsWithStream 实时处理结果流
func (r *RunnerWithCtx) HandleResultsWithStream(ctx context.Context, handler JobRetHandlerWithCtx) {
	r.mu.Lock()
	if r.isHandled {
		r.mu.Unlock()
		return
	}
	r.isHandled = true
	r.mu.Unlock()

	// 实时监听结果，直到所有任务完成或上下文被取消
	for {
		select {
		case ret, ok := <-r.results:
			if !ok {
				// 结果通道已关闭
				return
			}
			handler(ctx, ret)
		case <-r.ctx.Done():
			// 上下文被取消，停止处理
			return
		}
	}
}

// HandleResultsWithStreamAndOutput 实时处理结果流并传递输出参数
func (r *RunnerWithCtx) HandleResultsWithStreamAndOutput(ctx context.Context, handler JobRetOutputHandlerWithCtx, output any) {
	r.mu.Lock()
	if r.isHandled {
		r.mu.Unlock()
		return
	}
	r.isHandled = true
	r.mu.Unlock()

	// 实时监听结果，直到所有任务完成或上下文被取消
	for {
		select {
		case ret, ok := <-r.results:
			if !ok {
				// 结果通道已关闭
				return
			}
			handler(ctx, ret, output)
		case <-r.ctx.Done():
			// 上下文被取消，停止处理
			return
		}
	}
}

// HandleAllResultsWith 等待所有任务完成后处理结果
func (r *RunnerWithCtx) HandleAllResultsWith(ctx context.Context, handler JobRetHandlerWithCtx) {
	r.mu.Lock()
	if r.isHandled {
		r.mu.Unlock()
		return
	}
	r.isHandled = true
	r.mu.Unlock()

	// 创建一个goroutine来等待WaitGroup
	done := make(chan struct{})
	go func() {
		r.wg.Wait()
		close(done)
	}()

	// 等待任务完成或上下文被取消
	select {
	case <-done:
		// 所有任务完成
		close(r.results)
		for ret := range r.results {
			handler(ctx, ret)
		}
	case <-r.ctx.Done():
		// 上下文被取消
		return
	}
}

// Cancel 取消所有正在运行的任务
func (r *RunnerWithCtx) Cancel() {
	r.cancel()
}

// Context 获取Runner的上下文
func (r *RunnerWithCtx) Context() context.Context {
	return r.ctx
}

// GetJobsCount 获取任务总数
func (r *RunnerWithCtx) GetJobsCount() int {
	return r.jobsCount
}

// GetJobStatus 获取指定任务的状态
func (r *RunnerWithCtx) GetJobStatus(jobID string) (int, bool) {
	if value, ok := r.jobs.Load(jobID); ok {
		job := value.(*JobWithCtx)
		return job.RunStatus, true
	}
	return -1, false
}
