# 拼音工具

这是一个Go语言汉字转拼音工具包，提供了将中文字符转换为拼音的功能，支持自定义分隔符和数组返回格式。

## 功能特点

- 将中文字符转换为拼音字符串
- 将中文字符转换为拼音数组
- 支持自定义拼音之间的分隔符
- 保留非中文字符
- 智能识别汉字字符（Unicode范围：0x4E00~0x9FA5）
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/pinyin_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/pinyin_tools"
```

### 将汉字转换为拼音字符串

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/pinyin_tools"
)

func main() {
    // 不使用分隔符
    result := pinyin_tools.ConvertToPinyin("你好世界", "")
    fmt.Println(result) // 输出: nihaoshijie
    
    // 使用连字符作为分隔符
    result = pinyin_tools.ConvertToPinyin("你好世界", "-")
    fmt.Println(result) // 输出: ni-hao-shi-jie
    
    // 处理混合中英文字符
    result = pinyin_tools.ConvertToPinyin("Hello你好World世界", " ")
    fmt.Println(result) // 输出: H e l l o ni hao W o r l d shi jie
}
```

### 将汉字转换为拼音数组

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/pinyin_tools"
)

func main() {
    // 获取拼音数组
    result := pinyin_tools.ConvertToPinyinList("你好s世界")
    fmt.Println(result) // 输出: [ni hao s shi jie]
    
    // 处理混合中英文字符
    result = pinyin_tools.ConvertToPinyinList("Hello你好123")
    fmt.Println(result) // 输出: [H e l l o ni hao 1 2 3]
}
```

### 处理包含非中文字符的文本

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/pinyin_tools"
)

func main() {
    // 包含数字和标点符号
    result := pinyin_tools.ConvertToPinyin("你好2023！世界", "_")
    fmt.Println(result) // 输出: ni_hao_2_0_2_3_！_shi_jie
    
    // 包含英文字母
    result = pinyin_tools.ConvertToPinyin("中文English混合", "-")
    fmt.Println(result) // 输出: zhong-wen-E-n-g-l-i-s-h-hun-he
}
```

## API 文档

### ConvertToPinyin 函数

```go
func ConvertToPinyin(str, split string) string
```

将汉字转换为拼音字符串。

**参数：**
- `str`: 待转换的中文字符串
- `split`: 拼音之间的分隔符

**返回值：**
- 转换后的拼音字符串

**示例：**
```go
result := pinyin_tools.ConvertToPinyin("你好世界", "-")
// 返回: "ni-hao-shi-jie"
```

### ConvertToPinyinList 函数

```go
func ConvertToPinyinList(str string) []string
```

将汉字转换为拼音数组，每个元素对应原字符串中的一个字符。

**参数：**
- `str`: 待转换的中文字符串

**返回值：**
- 转换后的拼音字符串数组

**示例：**
```go
result := pinyin_tools.ConvertToPinyinList("你好世界")
// 返回: ["ni", "hao", "shi", "jie"]
```

### isChinese 函数（内部函数）

```go
func isChinese(c rune) bool
```

判断字符是否为汉字。这是一个内部辅助函数，用于识别Unicode范围为0x4E00~0x9FA5的汉字字符。

**参数：**
- `c`: 待判断的字符（rune类型）

**返回值：**
- 如果是汉字返回true，否则返回false

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