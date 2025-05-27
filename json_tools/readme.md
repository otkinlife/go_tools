# JSON 工具

这是一个Go语言JSON处理工具包，提供了从字符串中提取JSON子串等功能，简化JSON数据处理。

## 功能特点

- 从文本字符串中提取JSON子串
- 支持嵌套的JSON对象和数组
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

### 从字符串中提取JSON

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

func main() {
    // 包含JSON的字符串
    str := "这是一段文本，其中包含JSON数据：{\"name\":\"John\",\"age\":30,\"city\":\"New York\"}"
    
    // 提取JSON子串
    jsonStr, err := json_tools.ExtractJsonFromStr(str)
    if err != nil {
        fmt.Println("提取JSON失败:", err)
        return
    }
    
    fmt.Println("提取的JSON:", jsonStr)
    // 输出: {"name":"John","age":30,"city":"New York"}
}
```

### 处理嵌套的JSON

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

func main() {
    // 包含嵌套JSON的字符串
    str := "用户提交的数据是 {\"user\":{\"name\":\"John\",\"details\":{\"age\":30,\"address\":{\"city\":\"New York\"}}}}"
    
    // 提取JSON子串
    jsonStr, err := json_tools.ExtractJsonFromStr(str)
    if err != nil {
        fmt.Println("提取JSON失败:", err)
        return
    }
    
    fmt.Println("提取的JSON:", jsonStr)
    // 输出: {"user":{"name":"John","details":{"age":30,"address":{"city":"New York"}}}}
}
```

### 处理JSON数组

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/json_tools"
)

func main() {
    // 包含JSON数组的字符串
    str := "系统返回的结果列表：[{\"id\":1,\"name\":\"Item 1\"},{\"id\":2,\"name\":\"Item 2\"}]"
    
    // 提取JSON子串
    jsonStr, err := json_tools.ExtractJsonFromStr(str)
    if err != nil {
        fmt.Println("提取JSON失败:", err)
        return
    }
    
    fmt.Println("提取的JSON:", jsonStr)
    // 输出: [{"id":1,"name":"Item 1"},{"id":2,"name":"Item 2"}]
}
```

## 参数说明

### ExtractJsonFromStr 函数

```go
func ExtractJsonFromStr(input string) (string, error)
```

- `input`: 包含JSON子串的输入字符串
- 返回值:
    - 提取出的JSON字符串
    - 错误信息（如果有）

## 错误处理

函数可能返回以下错误：

- `"no JSON found in input string"`: 输入字符串中没有找到JSON对象或数组
- `"incomplete JSON string"`: 输入字符串中的JSON不完整（括号不匹配）

## 注意事项

- 函数会提取第一个有效的JSON对象或数组
- 支持嵌套的JSON结构
- 函数通过匹配花括号`{}`和方括号`[]`来识别JSON，确保输入字符串中的JSON格式正确