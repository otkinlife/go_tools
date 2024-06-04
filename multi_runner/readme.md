# Multi-Runner 使用手册

## 目录
- [简介](#简介)
- [安装](#安装)
- [快速开始](#快速开始)
- [API 详解](#api-详解)
    - [NewRunner](#newrunner)
    - [AddJob](#addjob)
    - [Run](#run)
    - [HandleResultsWithStream](#handleresultswithstream)
    - [HandleAllResultsWith](#handleallresultswith)
- [示例](#示例)
    - [基本示例](#基本示例)
    - [处理结果示例](#处理结果示例)
- [注意事项](#注意事项)

## 简介

`Multi-Runner` 是一个用于并发执行任务的 Go 包。它允许用户添加多个任务，并以并发的方式执行这些任务，并提供了处理任务结果的方式。

## 安装

可以通过以下命令安装 `Multi-Runner` 包：

```bash
go get github.com/otkinlife/go_tools/multi_runner
```

## 快速开始

下面是一个快速开始的示例：

```go
package main

import (
	"fmt"
	"github.com/yourusername/multi_runner"
)

func main() {
	runner := multi_runner.NewRunner(5)

	// 添加任务
	for i := 0; i < 10; i++ {
		runner.AddJob(func(data any) multi_runner.JobRet {
			return multi_runner.JobRet{Data: data}
		}, i)
	}

	// 运行任务
	runner.Run()

	// 处理结果
	runner.HandleResultsWithStream(func(ret multi_runner.JobRet) {
		if ret.Err != nil {
			fmt.Printf("任务出错: %v\n", ret.Err)
		} else {
			fmt.Printf("任务结果: %v\n", ret.Data)
		}
	})
}
```

## API 详解

### NewRunner

创建一个新的 `Runner` 实例。

```go
func NewRunner(maxSize int) *Runner
```

- `maxSize`: 最大同时并发任务数。

### AddJob

向 `Runner` 添加一个任务。

```go
func (r *Runner) AddJob(handler JobExecute, runParams any) error
```

- `handler`: 任务执行函数。
- `runParams`: 任务执行参数。

### Run

开始执行所有添加的任务。

```go
func (r *Runner) Run()
```

### HandleResultsWithStream

以流的方式处理任务结果。

```go
func (r *Runner) HandleResultsWithStream(handler JobRetHandler)
```

- `handler`: 结果处理函数。

### HandleAllResultsWith

等待所有任务完成后处理结果。

```go
func (r *Runner) HandleAllResultsWith(handler JobRetHandler)
```

- `handler`: 结果处理函数。

## 示例

### 基本示例

下面是一个基本示例，展示了如何创建 Runner，添加任务，并执行任务。

```go
package main

import (
	"fmt"
	"github.com/yourusername/multi_runner"
)

func main() {
	runner := multi_runner.NewRunner(5)

	// 添加任务
	for i := 0; i < 10; i++ {
		runner.AddJob(func(data any) multi_runner.JobRet {
			return multi_runner.JobRet{Data: data}
		}, i)
	}

	// 运行任务
	runner.Run()

	// 处理结果
	runner.HandleResultsWithStream(func(ret multi_runner.JobRet) {
		if ret.Err != nil {
			fmt.Printf("任务出错: %v\n", ret.Err)
		} else {
			fmt.Printf("任务结果: %v\n", ret.Data)
		}
	})
}
```

### 处理结果示例

下面是一个示例，展示了如何等待所有任务完成后再处理结果。

```go
package main

import (
	"fmt"
	"github.com/yourusername/multi_runner"
)

func main() {
	runner := multi_runner.NewRunner(5)

	// 添加任务
	for i := 0; i < 10; i++ {
		runner.AddJob(func(data any) multi_runner.JobRet {
			return multi_runner.JobRet{Data: data}
		}, i)
	}

	// 运行任务
	runner.Run()

	// 处理所有结果
	runner.HandleAllResultsWith(func(ret multi_runner.JobRet) {
		if ret.Err != nil {
			fmt.Printf("任务出错: %v\n", ret.Err)
		} else {
			fmt.Printf("任务结果: %v\n", ret.Data)
		}
	})
}
```

## 注意事项

- 确保在调用 `Run` 方法之前已经添加了所有任务。
- `HandleResultsWithStream` 和 `HandleAllResultsWith` 方法只能调用一个，并且只能调用一次。
- 并发环境下需要注意对共享资源的访问保护，例如 `jobs` 和 `isHandled` 标志位。

通过以上的使用手册，你应该能够快速上手 `Multi-Runner` 并在你的项目中使用它来并发执行任务。