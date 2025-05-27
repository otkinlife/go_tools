# 文件下载工具

这是一个Go语言文件下载工具包，提供了文件下载和图片下载的功能，支持分块下载大文件和图片格式检测。

## 功能特点

- 支持单线程和多线程分块下载
- 自动检测服务器是否支持Range请求
- 支持图片专用下载功能，自动检测图片格式
- 支持自定义保存路径和文件名
- 完整的错误处理

## 安装

```bash
go get github.com/otkinlife/go_tools/downloader
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/downloader"
```

### 下载普通文件

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/downloader"
)

func main() {
    // 下载文件，不使用分块下载
    err := downloader.DownloadFile("https://example.com/file.zip", "./downloads/myfile.zip", 0)
    if err != nil {
        fmt.Println("下载失败:", err)
        return
    }
    
    fmt.Println("文件下载成功")
    
    // 使用12个块并行下载大文件
    err = downloader.DownloadFile("https://example.com/largefile.zip", "./downloads/", 12)
    if err != nil {
        fmt.Println("下载失败:", err)
        return
    }
    
    fmt.Println("大文件下载成功")
}
```

### 下载图片

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/downloader"
)

func main() {
    // 下载图片，自动检测格式并生成MD5文件名
    result := downloader.DownloadImage("https://example.com/image.jpg", "./images/")
    if result.Err != nil {
        fmt.Println("图片下载失败:", result.Err)
        return
    }
    
    fmt.Println("图片下载成功:")
    fmt.Println("- 保存目录:", result.Dir)
    fmt.Println("- 文件名:", result.FileName)
    fmt.Println("- 完整路径:", result.Filepath)
    fmt.Println("- 图片格式:", result.Format)
}
```

## 参数说明

### DownloadFile 函数

```go
func DownloadFile(url string, filePath string, chunkCount int) error
```

- `url`: 要下载的文件URL
- `filePath`: 保存文件的路径
    - 如果以`/`或`\`结尾，则视为目录，文件名将从URL中提取
    - 否则视为完整的文件路径（包含文件名）
- `chunkCount`: 分块下载的块数
    - `0`: 不使用分块下载
    - `1-32`: 使用指定数量的块并行下载
    - 超过32将返回错误

### DownloadImage 函数

```go
func DownloadImage(url, targetDir string) DImgRet
```

- `url`: 要下载的图片URL
- `targetDir`: 保存图片的目录
    - 如果为空，则使用系统临时目录

返回`DImgRet`结构体，包含以下字段：
- `Dir`: 图片保存的目录
- `FileName`: 图片文件名（使用图片内容的MD5值作为文件名）
- `Filepath`: 完整的文件路径
- `Err`: 错误信息（如果有）
- `Format`: 图片格式（如"jpeg"、"png"等）

## 注意事项

- 分块下载功能要求服务器支持Range请求，如果不支持会自动回退到单线程下载
- 图片下载功能会自动检测图片格式，不支持的格式会返回错误
- 文件下载过程中会创建临时文件（文件名后缀为`.download`），下载完成后重命名
- 最大支持32个并行下载块
