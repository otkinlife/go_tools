# Go Trace 包装器

这个包提供了一个简单的包装器，用于封装Go语言内建的 `trace` 包。这个包可以用于追踪Go程序的执行，对于调试和性能优化非常有用。

## 安装

要安装这个包，你可以使用 `go get` 命令：

```bash
go get github.com/otkinlife/go_tools/trace
```

## 使用

首先，导入包：

```go
import "github.com/otkinlife/go_tools/trace"
```

然后，你可以创建一个新的 `Trace` 实例，并开始追踪：

```go
func main() {
	t := trace.New("trace.out")
	err := t.Start()
	if err != nil {
		panic(err)
	}
	defer t.Stop()

	// 你的程序在这里
}
```

在这个例子中，我们首先创建了一个新的 `Trace` 实例，并开始追踪。然后，在程序结束时，我们调用 `Stop` 方法来停止追踪。所有的追踪信息都会被写入到 "trace.out" 文件中。

## 分析追踪数据
```shell
go tool trace trace.out 
```
[报告使用方式](./report_usage.md)
## 注意事项

`trace` 包主要用于调试和性能优化，在生产环境中通常不会使用，因为它会增加一些运行开销。