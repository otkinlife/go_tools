# 数字工具

这是一个Go语言数字处理工具包，提供了数字格式化和转换功能，简化数值处理操作。

## 功能特点

- 支持泛型，适用于各种数值类型
- 数字格式化，添加千分位分隔符
- 整数转罗马数字
- 支持自定义分隔符和小数位数
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/number_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/number_tools"
```

### 数字格式化（添加千分位分隔符）

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/number_tools"
)

func main() {
    // 使用逗号作为千分位分隔符，保留2位小数
    formatted := number_tools.FormatNumberWithCommas(1234567.89, ",", 2)
    fmt.Println(formatted) // 输出: 1,234,567.89
    
    // 使用点作为千分位分隔符，保留3位小数
    formatted = number_tools.FormatNumberWithCommas(1234567.89, ".", 3)
    fmt.Println(formatted) // 输出: 1.234.567.890
    
    // 不保留小数位
    formatted = number_tools.FormatNumberWithCommas(1234567, ",", 0)
    fmt.Println(formatted) // 输出: 1,234,567
    
    // 支持各种数值类型
    formatted = number_tools.FormatNumberWithCommas(int64(9876543210), ",", 0)
    fmt.Println(formatted) // 输出: 9,876,543,210
    
    formatted = number_tools.FormatNumberWithCommas(float32(12345.6), ",", 1)
    fmt.Println(formatted) // 输出: 12,345.6
}
```

### 整数转罗马数字

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/number_tools"
)

func main() {
    // 将整数转换为罗马数字
    roman := number_tools.IntToRoman(3)
    fmt.Println("3 的罗马数字表示:", roman) // 输出: III
    
    roman = number_tools.IntToRoman(9)
    fmt.Println("9 的罗马数字表示:", roman) // 输出: IX
    
    roman = number_tools.IntToRoman(40)
    fmt.Println("40 的罗马数字表示:", roman) // 输出: XL
    
    roman = number_tools.IntToRoman(1984)
    fmt.Println("1984 的罗马数字表示:", roman) // 输出: MCMLXXXIV
}
```

## 函数说明

### FormatNumberWithCommas 函数

```go
func FormatNumberWithCommas[T Number](originalNum T, split string, decimalPlaces int) string
```

- `originalNum`: 要格式化的数字（支持各种数值类型）
- `split`: 千分位分隔符（如","、"."等）
- `decimalPlaces`: 保留的小数位数
- 返回值: 格式化后的数字字符串

### IntToRoman 函数

```go
func IntToRoman(num int) string
```

- `num`: 要转换的整数（范围：1-3999）
- 返回值: 对应的罗马数字字符串

## 支持的数值类型

`FormatNumberWithCommas` 函数通过泛型支持以下数值类型：

- 整数类型：int, int8, int16, int32, int64
- 无符号整数类型：uint, uint8, uint16, uint32, uint64
- 浮点数类型：float32, float64

## 注意事项

- `IntToRoman` 函数适用于1到3999之间的整数，超出此范围的值可能导致不正确的结果
- `FormatNumberWithCommas` 函数会根据指定的小数位数进行四舍五入
- 使用泛型功能需要Go 1.18或更高版本