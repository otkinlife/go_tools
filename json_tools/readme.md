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

所有函数都提供了完善的错误处理机制：

- 序列化失败时返回详细的错误信息
- 反序列化失败时提供错误原因
- 使用 `fmt.Errorf` 包装错误，提供更多上下文信息

## 注意事项

- 所有函数都支持任意Go数据类型的处理
- JSON验证函数性能优化，适合高频调用
- 格式化输出使用2个空格作为缩进
- 错误信息采用中文描述，便于调试