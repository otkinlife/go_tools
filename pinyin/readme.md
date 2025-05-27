# 拼音工具

这是一个Go语言汉字转拼音工具包，提供了将中文字符转换为拼音的功能，支持自定义分隔符。

## 功能特点

- 将中文字符转换为拼音
- 支持自定义拼音之间的分隔符
- 保留非中文字符
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/pinyin
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/pinyin"
```

### 将汉字转换为拼音

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/pinyin"
)

func main() {
    // 不使用分隔符
    result := pinyin.ConvertToPinyin("你好世界", "")
    fmt.Println(result) // 输出: nihaoshijie
    
    // 使用连字符作为分隔符
    result = pinyin.ConvertToPinyin("你好世界", "-")
    fmt.Println(result) // 输出: ni-hao-shi-jie
    
    // 处理混合中英文字符
    result = pinyin.ConvertToPinyin("Hello你好World世界", " ")
    fmt.Println(result) // 输出: H e l l o ni hao W o r l d shi jie
}
```

### 处理包含非中文字符的文本

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/pinyin"
)

func main() {
    // 包含数字和标点符号
    result := pinyin.ConvertToPinyin("你好2023！世界", "_")
    fmt.Println(result) // 输出: ni_hao_2_0_2_3_！_shi_jie
    
    // 包含英文字母
    result = pinyin.ConvertToPinyin("中文English混合", "-")
    fmt.Println(result) // 输出: zhong-wen-E-n-g-l-i-s-h-hun-he
}
```

## 函数说明

### ConvertToPinyin 函数

```go
func ConvertToPinyin(str, split string) string
```

- `str`: 待转换的中文字符串
- `split`: 拼音之间的分隔符
- 返回值: 转换后的拼音字符串

## 工作原理

1. 使用 [github.com/mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin) 库将中文字符转换为拼音
2. 遍历原始字符串中的每个字符：
    - 如果是中文字符（Unicode范围：0x4E00~0x9FA5），则替换为对应的拼音
    - 如果不是中文字符，则保持原样
3. 使用指定的分隔符连接所有字符

## 注意事项

- 非中文字符会被保留在结果字符串中
- 分隔符会应用于所有字符之间，包括非中文字符
- 默认使用小写拼音，不包含声调