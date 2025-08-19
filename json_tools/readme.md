# JSON 工具

这是一个Go语言JSON处理工具包，提供了JSON序列化、反序列化、格式化、验证等功能，简化JSON数据处理。

## 功能特点

- JSON字符串解析和生成
- JSON数据序列化和反序列化
- JSON格式化输出（美化）
- JSON有效性验证
- 完善的错误处理机制
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/json_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/json_tools"
```

### JSON序列化

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func main() {
    user := User{
        Name: "John",
        Age:  30,
        City: "New York",
    }
    
    // 序列化为JSON字符串
    jsonStr, err := json_tools.MarshalJson(user)
    if err != nil {
        fmt.Println("序列化失败:", err)
        return
    }
    
    fmt.Println("JSON字符串:", jsonStr)
    // 输出: {"name":"John","age":30,"city":"New York"}
}
```

### JSON反序列化

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func main() {
    jsonStr := `{"name":"John","age":30,"city":"New York"}`
    
    var user User
    err := json_tools.UnmarshalJson(jsonStr, &user)
    if err != nil {
        fmt.Println("反序列化失败:", err)
        return
    }
    
    fmt.Printf("用户信息: %+v\n", user)
    // 输出: 用户信息: {Name:John Age:30 City:New York}
}
```

### JSON格式化输出

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

func main() {
    data := map[string]interface{}{
        "name": "John",
        "age":  30,
        "address": map[string]string{
            "city":    "New York",
            "country": "USA",
        },
    }
    
    // 格式化输出JSON
    prettyJson, err := json_tools.MarshalJsonPretty(data)
    if err != nil {
        fmt.Println("格式化失败:", err)
        return
    }
    
    fmt.Println("格式化的JSON:")
    fmt.Println(prettyJson)
}
```

### JSON有效性验证

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

func main() {
    validJson := `{"name":"John","age":30}`
    invalidJson := `{"name":"John","age":}`
    
    fmt.Println("有效JSON:", json_tools.IsValidJson(validJson))     // true
    fmt.Println("无效JSON:", json_tools.IsValidJson(invalidJson))   // false
}
```

### 无错误返回的JSON操作

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func main() {
    user := User{
        Name: "John",
        Age:  30,
        City: "New York",
    }
    
    // 序列化（无错误返回，错误会记录到日志）
    jsonStr := json_tools.MarshalWTE(user)
    fmt.Println("JSON字符串:", jsonStr)
    
    // 反序列化（无错误返回，错误会记录到日志）
    var newUser User
    json_tools.UnmarshalWTE(jsonStr, &newUser)
    fmt.Printf("用户信息: %+v\n", newUser)
}
```

## 函数说明

### UnmarshalJson 函数

```go
func UnmarshalJson(data string, v any) error
```

- `data`: 需要解析的JSON字符串
- `v`: 接收解析结果的变量指针
- 返回值: 错误信息（如果有）

### MarshalJson 函数

```go
func MarshalJson(v any) (string, error)
```

- `v`: 需要序列化的Go变量
- 返回值:
    - JSON字符串
    - 错误信息（如果有）

### UnmarshalWTE 函数

```go
func UnmarshalWTE(data string, v any)
```

- `data`: 需要解析的JSON字符串
- `v`: 接收解析结果的变量指针
- 注意: 不返回错误，解析失败时会记录错误日志

### MarshalWTE 函数

```go
func MarshalWTE(v any) string
```

- `v`: 需要序列化的Go变量
- 返回值: JSON字符串（失败时返回空字符串）
- 注意: 不返回错误，序列化失败时会记录错误日志并返回空字符串

### MarshalJsonPretty 函数

```go
func MarshalJsonPretty(v any) (string, error)
```

- `v`: 需要序列化的Go变量
- 返回值:
    - 格式化的JSON字符串（带缩进）
    - 错误信息（如果有）

### IsValidJson 函数

```go
func IsValidJson(data string) bool
```

- `data`: 需要验证的JSON字符串
- 返回值: 布尔值，表示JSON是否有效

## 错误处理

工具包提供两种错误处理方式：

### 标准错误处理
- `UnmarshalJson` 和 `MarshalJson` 等函数返回详细的错误信息
- 使用 `fmt.Errorf` 包装错误，提供更多上下文信息
- 适合需要精确错误处理的场景

### WTE（Without Throwing Error）模式
- `UnmarshalWTE` 和 `MarshalWTE` 函数不返回错误
- 错误会自动记录到日志系统中
- 适合对错误处理要求不严格的快速开发场景
- 依赖 `logger_tools` 包进行错误日志记录

## 注意事项

- 所有函数都支持任意Go数据类型的处理
- JSON验证函数性能优化，适合高频调用
- 格式化输出使用2个空格作为缩进
- WTE模式的函数需要确保已正确配置日志系统
- 建议在生产环境中使用标准错误处理模式以便于调试