# 图片工具

这是一个Go语言图片处理工具包，提供了图片Base64编码/解码、格式判断和图片切割等功能。

## 功能特点

- 图片与Base64字符串的相互转换
- 图片格式判断
- 图片等分切割功能
- 支持从URL下载图片并切割

## 安装

```bash
go get github.com/otkinlife/go_tools/img
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/img"
```

### 图片转Base64

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/img"
)

func main() {
    // 将图片转换为Base64字符串
    base64Str, err := img.EncodeImg2Base64Str("path/to/image.jpg")
    if err != nil {
        fmt.Println("转换失败:", err)
        return
    }
    
    fmt.Println("Base64字符串:", base64Str)
}
```

### Base64转图片

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/img"
)

func main() {
    // Base64字符串（格式：data:image/png;base64,xxxxx）
    base64Str := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA..."
    
    // 将Base64字符串转换为图片并保存
    err := img.DecodeBase642Img(base64Str, "output.png")
    if err != nil {
        fmt.Println("转换失败:", err)
        return
    }
    
    fmt.Println("图片已保存")
}
```

### 判断文件是否为图片格式

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/img"
)

func main() {
    // 判断文件后缀是否为图片格式
    isImage := img.IsImage("jpg")
    fmt.Println("是否为图片格式:", isImage) // 输出: true
    
    isImage = img.IsImage("txt")
    fmt.Println("是否为图片格式:", isImage) // 输出: false
}
```

### 图片等分切割

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/img"
)

func main() {
    // 将图片平均切割成9份（3x3网格）
    imagePaths, err := img.AvgSplitImg(9, "input.png", "./output/")
    if err != nil {
        fmt.Println("切割失败:", err)
        return
    }
    
    fmt.Println("切割后的图片路径:", imagePaths)
}
```

### 从URL下载并切割图片

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/img"
)

func main() {
    // 从URL下载图片并切割成4份（2x2网格）
    imagePaths, err := img.AvgSplitImgFromUrl(4, "https://example.com/image.jpg", "./output/")
    if err != nil {
        fmt.Println("下载或切割失败:", err)
        return
    }
    
    fmt.Println("切割后的图片路径:", imagePaths)
}
```

## 支持的图片格式

工具包支持以下图片格式：
- JPG/JPEG
- PNG
- GIF
- BMP
- WEBP

## 注意事项

- 图片切割功能要求切割份数必须是完全平方数（如1, 4, 9, 16等）
- Base64字符串必须包含前缀（如`data:image/png;base64,`）
