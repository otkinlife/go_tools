# CSV 工具

这是一个 Go 语言 CSV 处理工具包，提供了简单易用的 CSV 文件读取、操作和保存功能。

## 功能特点

- 从文件、字符串或二维数组加载 CSV 数据
- 支持自定义分隔符
- 通过列名访问数据
- 支持添加和修改数据
- 保存 CSV 数据到文件
- 完整的错误处理

## 安装

```bash
go get github.com/otkinlife/go_tools/csv_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/csv_tools"
```

### 创建 CSV 数据对象

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    // 创建新的 CSV 数据对象
    csvData := csv_tools.NewCSVData()
}
```

### 从字符串加载 CSV 数据

```go
package main

import (
    "fmt"
    "strings"
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    csvContent := "姓名,年龄,邮箱\n张三,30,zhangsan@example.com\n李四,25,lisi@example.com"
    reader := strings.NewReader(csvContent)
    
    csvData := csv_tools.NewCSVData()
    csvData.LoadFromReader(reader)
    
    // 检查是否有错误
    if err := csvData.GetError(); err != nil {
        fmt.Println("加载 CSV 出错:", err)
        return
    }
    
    // 获取数据
    data := csvData.GetData()
    fmt.Println("CSV 数据:", data)
}
```

### 使用自定义分隔符

```go
package main

import (
    "fmt"
    "strings"
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    // 使用分号作为分隔符
    csvContent := "姓名;年龄;邮箱\n张三;30;zhangsan@example.com\n李四;25;lisi@example.com"
    reader := strings.NewReader(csvContent)
    
    csvData := csv_tools.NewCSVData().SetSplit(";")
    csvData.LoadFromReader(reader)
    
    // 获取数据
    data := csvData.GetData()
    fmt.Println("CSV 数据:", data)
}
```

### 从二维数组加载数据

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    lines := [][]string{
        {"姓名", "年龄", "邮箱"},
        {"张三", "30", "zhangsan@example.com"},
        {"李四", "25", "lisi@example.com"},
    }
    
    csvData := csv_tools.NewCSVData()
    csvData.LoadFromLines(lines)
    
    // 获取表头
    headers := csvData.GetHeaders()
    fmt.Println("表头映射:", headers)
}
```

### 访问特定数据

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    lines := [][]string{
        {"姓名", "年龄", "邮箱"},
        {"张三", "30", "zhangsan@example.com"},
        {"李四", "25", "lisi@example.com"},
    }
    
    csvData := csv_tools.NewCSVData()
    csvData.LoadFromLines(lines)
    
    // 获取特定列的索引
    ageIndex := csvData.GetHeaderIndex("年龄")
    fmt.Println("年龄列的索引:", ageIndex)
    
    // 获取特定行列的值
    email := csvData.GetLineValue("邮箱", 0) // 第一行的邮箱
    fmt.Println("张三的邮箱:", email)
}
```

### 添加数据

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    csvData := csv_tools.NewCSVData()
    
    // 设置表头
    csvData.SetHeaders([]string{"姓名", "年龄", "邮箱"})
    
    // 添加数据行
    csvData.AppendData([]string{"张三", "30", "zhangsan@example.com"})
    csvData.AppendData([]string{"李四", "25", "lisi@example.com"})
    
    // 获取所有数据
    data := csvData.GetData()
    fmt.Println("CSV 数据:", data)
}
```

### 保存到文件

```go
package main

import (
    "github.com/otkinlife/go_tools/csv_tools"
)

func main() {
    csvData := csv_tools.NewCSVData()
    
    // 设置表头和数据
    csvData.SetHeaders([]string{"姓名", "年龄", "邮箱"})
    csvData.AppendData([]string{"张三", "30", "zhangsan@example.com"})
    csvData.AppendData([]string{"李四", "25", "lisi@example.com"})
    
    // 保存到文件
    csvData.SaveToFile("./output.csv")
    
    // 检查是否有错误
    if err := csvData.GetError(); err != nil {
        panic(err)
    }
}
```

## 错误处理

所有操作都会在内部记录错误，可以通过 `GetError()` 方法获取：

```go
csvData := csv_tools.NewCSVData()
// 执行各种操作...

if err := csvData.GetError(); err != nil {
    fmt.Println("发生错误:", err)
    return
}
```

## 测试

包中包含完整的测试用例，可以通过以下命令运行测试：

```bash
go test github.com/otkinlife/go_tools/csv_tools
```