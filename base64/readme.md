# Base64 工具

这是一个简单的 Go 语言 Base64 编码解码工具包，提供了字符串的 Base64 加密和解密功能。

## 功能

- Base64 编码：将普通字符串转换为 Base64 编码格式
- Base64 解码：将 Base64 编码字符串转换回普通字符串

## 安装

```bash
go get github.com/otkinlife/go_tools/base64
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/base64"
```

### Base64 编码

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/base64"
)

func main() {
    // 对字符串进行 Base64 编码
    encoded := base64.Encode("hello world")
    fmt.Println("编码结果:", encoded) // 输出: aGVsbG8gd29ybGQ=
}
```

### Base64 解码

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/base64"
)

func main() {
    // 对 Base64 编码的字符串进行解码
    decoded, err := base64.Decode("aGVsbG8gd29ybGQ=")
    if err != nil {
        fmt.Println("解码错误:", err)
        return
    }
    fmt.Println("解码结果:", decoded) // 输出: hello world
}
```

## 错误处理

解码函数会返回可能的错误，请确保在使用时进行错误检查：

```go
decoded, err := base64.Decode("invalid-base64-string")
if err != nil {
    // 处理错误
    fmt.Println("解码失败:", err)
}
```

## 测试

包中包含测试用例，可以通过以下命令运行测试：

```bash
go test github.com/otkinlife/go_tools/base64
```