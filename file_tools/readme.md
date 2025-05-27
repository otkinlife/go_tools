# 文件工具

这是一个Go语言文件操作工具包，提供了文件存在性检查等功能，简化文件系统操作。

## 功能特点

- 检查文件或目录是否存在
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/file_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/file_tools"
```

### 检查文件是否存在

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/file_tools"
)

func main() {
    // 检查文件是否存在
    exists := file_tools.IsFileExist("config.json")
    
    if exists {
        fmt.Println("文件存在")
    } else {
        fmt.Println("文件不存在")
    }
    
    // 检查目录是否存在
    dirExists := file_tools.IsFileExist("./logs")
    
    if dirExists {
        fmt.Println("目录存在")
    } else {
        fmt.Println("目录不存在")
    }
}
```

## 参数说明

### IsFileExist 函数

```go
func IsFileExist(filename string) bool
```

- `filename`: 要检查的文件或目录路径
- 返回值: 布尔值，表示文件或目录是否存在
    - `true`: 文件或目录存在
    - `false`: 文件或目录不存在

## 注意事项

- 此函数可用于检查文件和目录是否存在
- 如果文件存在但无法访问（如权限问题），函数仍可能返回`true`
- 路径可以是相对路径或绝对路径
