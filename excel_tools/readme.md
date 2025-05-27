# Excel 工具

这是一个Go语言Excel处理工具包，提供了Excel单元格地址转换等功能，简化Excel文件操作。

## 功能特点

- Excel单元格地址与行列号互相转换
- 支持标准Excel列名格式（如A、B、AA、AB等）
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/excel_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/excel_tools"
```

### 行列号转单元格地址

将行号和列号转换为Excel单元格地址（如：1行2列 → B1）

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/excel_tools"
)

func main() {
    // 获取第1行第2列的单元格地址
    cellAddress := excel_tools.GetExcelCellAddress(1, 2)
    fmt.Println("单元格地址:", cellAddress) // 输出: B1
    
    // 获取第5行第27列的单元格地址
    cellAddress = excel_tools.GetExcelCellAddress(5, 27)
    fmt.Println("单元格地址:", cellAddress) // 输出: AA5
}
```

### 单元格地址转行列号

将Excel单元格地址转换为行号和列号（如：B1 → 1行2列）

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/excel_tools"
)

func main() {
    // 获取B1单元格的行号和列号
    row, column := excel_tools.GetExcelCellRowColumn("B1")
    fmt.Printf("行号: %d, 列号: %d\n", row, column) // 输出: 行号: 1, 列号: 2
    
    // 获取AA5单元格的行号和列号
    row, column = excel_tools.GetExcelCellRowColumn("AA5")
    fmt.Printf("行号: %d, 列号: %d\n", row, column) // 输出: 行号: 5, 列号: 27
}
```

## 参数说明

### GetExcelCellAddress 函数

```go
func GetExcelCellAddress(row, column int) string
```

- `row`: 行号（从1开始）
- `column`: 列号（从1开始）
- 返回值: Excel单元格地址（如"A1"、"B2"等）

### GetExcelCellRowColumn 函数

```go
func GetExcelCellRowColumn(cellAddress string) (int, int)
```

- `cellAddress`: Excel单元格地址（如"A1"、"B2"等）
- 返回值: 行号（从1开始）, 列号（从1开始）
    - 如果单元格地址无效，则返回 -1, -1

## 注意事项

- 行号和列号都是从1开始计数的，符合Excel的习惯
- 列名支持多字母格式（如AA、AB等），最大可支持Excel的全部列范围
- 无效的单元格地址会返回错误值并记录日志
