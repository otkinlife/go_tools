# 国际化工具

这是一个Go语言国际化(i18n)工具包，提供了多语言文本管理和翻译功能，简化应用程序的国际化实现。

## 功能特点

- 支持多种语言的文本翻译
- 支持从JSON配置文件加载翻译内容
- 支持参数替换功能
- 内置常用语言代码常量
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/i18n_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/i18n_tools"
```

### 创建国际化工具实例

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/i18n_tools"
)

func main() {
    // 创建配置
    config := map[string]map[string]string{
        "hello": {
            i18n_tools.StandardLangEn:   "Hello",
            i18n_tools.StandardLangZhCN: "你好",
            i18n_tools.StandardLangJa:   "こんにちは",
        },
        "welcome": {
            i18n_tools.StandardLangEn:   "Welcome to our website",
            i18n_tools.StandardLangZhCN: "欢迎访问我们的网站",
            i18n_tools.StandardLangJa:   "私たちのウェブサイトへようこそ",
        },
    }
    
    // 创建国际化工具实例
    i18n := i18n_tools.NewI18NBuilder(config)
    
    // 获取翻译
    fmt.Println(i18n.Get("hello", i18n_tools.StandardLangEn))   // 输出: Hello
    fmt.Println(i18n.Get("hello", i18n_tools.StandardLangZhCN)) // 输出: 你好
    fmt.Println(i18n.Get("hello", i18n_tools.StandardLangJa))   // 输出: こんにちは
}
```

### 从JSON字符串创建国际化工具实例

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/i18n_tools"
)

func main() {
    // JSON配置
    jsonConfig := `{
        "hello": {
            "en": "Hello",
            "zh-CN": "你好",
            "ja": "こんにちは"
        },
        "welcome": {
            "en": "Welcome to our website",
            "zh-CN": "欢迎访问我们的网站",
            "ja": "私たちのウェブサイトへようこそ"
        }
    }`
    
    // 从JSON创建国际化工具实例
    i18n, err := i18n_tools.NewI18NBuilderFromJson(jsonConfig)
    if err != nil {
        fmt.Println("创建国际化工具失败:", err)
        return
    }
    
    // 获取翻译
    fmt.Println(i18n.Get("hello", "en"))     // 输出: Hello
    fmt.Println(i18n.Get("welcome", "zh-CN")) // 输出: 欢迎访问我们的网站
}
```

### 使用参数替换

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/i18n_tools"
)

func main() {
    // 创建配置
    config := map[string]map[string]string{
        "greeting": {
            i18n_tools.StandardLangEn:   "Hello, {name}! Today is {day}.",
            i18n_tools.StandardLangZhCN: "你好，{name}！今天是{day}。",
        },
    }
    
    // 创建国际化工具实例
    i18n := i18n_tools.NewI18NBuilder(config)
    
    // 定义参数
    params := map[string]any{
        "{name}": "John",
        "{day}":  "Monday",
    }
    
    // 获取带参数的翻译
    fmt.Println(i18n.GetWithParams("greeting", i18n_tools.StandardLangEn, params))
    // 输出: Hello, John! Today is Monday.
    
    fmt.Println(i18n.GetWithParams("greeting", i18n_tools.StandardLangZhCN, params))
    // 输出: 你好，John！今天是Monday。
}
```

### 获取特定键的所有语言配置

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/i18n_tools"
)

func main() {
    // 创建配置
    config := map[string]map[string]string{
        "hello": {
            i18n_tools.StandardLangEn:   "Hello",
            i18n_tools.StandardLangZhCN: "你好",
            i18n_tools.StandardLangJa:   "こんにちは",
        },
    }
    
    // 创建国际化工具实例
    i18n := i18n_tools.NewI18NBuilder(config)
    
    // 获取特定键的所有语言配置
    helloConfig := i18n.GetConfigWithKey("hello")
    fmt.Println(helloConfig) 
    // 输出: map[en:Hello ja:こんにちは zh-CN:你好]
}
```

### 获取特定语言的所有键值配置

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/i18n_tools"
)

func main() {
    // 创建配置
    config := map[string]map[string]string{
        "hello": {
            i18n_tools.StandardLangEn: "Hello",
            i18n_tools.StandardLangJa: "こんにちは",
        },
        "goodbye": {
            i18n_tools.StandardLangEn: "Goodbye",
            i18n_tools.StandardLangJa: "さようなら",
        },
    }
    
    // 创建国际化工具实例
    i18n := i18n_tools.NewI18NBuilder(config)
    
    // 获取特定语言的所有键值配置
    enConfig := i18n.GetKeyConfig(i18n_tools.StandardLangEn)
    fmt.Println(enConfig)
    // 输出: map[goodbye:goodbye hello:hello]
}
```

## 支持的语言

工具包内置了以下语言代码常量：

- `StandardLangEn` - 英语 (en)
- `StandardLangZhCN` - 简体中文 (zh-CN)
- `StandardLangZhTW` - 繁体中文 (zh-TW)
- `StandardLangJa` - 日语 (ja)
- `StandardLangKo` - 韩语 (ko)
- `StandardLangFr` - 法语 (fr)
- `StandardLangDe` - 德语 (de)
- `StandardLangEs` - 西班牙语 (es)
- `StandardLangIt` - 意大利语 (it)
- `StandardLangRu` - 俄语 (ru)
- 以及更多其他语言...

## 注意事项

- 如果请求的键或语言不存在，`Get`方法将返回键本身作为默认值
- 参数替换使用简单的字符串替换，确保参数名在文本中是唯一的
- 建议使用标准语言代码常量以确保一致性