# 列表工具

这是一个Go语言列表/切片处理工具包，提供了常用的列表操作功能，简化切片数据处理。

## 功能特点

- 支持泛型，适用于各种数据类型
- 判断元素是否在列表中
- 查找两个切片的差集
- 切片去重
- 切片分割成多个子切片
- 简单易用的API接口

## 安装

```bash
go get github.com/otkinlife/go_tools/list_tools
```

## 使用方法

### 导入包

```go
import "github.com/otkinlife/go_tools/list_tools"
```

### 判断元素是否在列表中

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/list_tools"
)

func main() {
    // 判断字符串是否在列表中
    result := list_tools.InList("apple", []string{"banana", "apple", "orange"})
    fmt.Println("apple 是否在列表中:", result) // 输出: true
    
    // 判断数字是否在列表中
    result = list_tools.InList(5, []int{1, 3, 5, 7, 9})
    fmt.Println("5 是否在列表中:", result) // 输出: true
    
    // 判断自定义类型是否在列表中
    type Person struct {
        Name string
        Age  int
    }
    
    p1 := Person{Name: "Alice", Age: 30}
    p2 := Person{Name: "Bob", Age: 25}
    p3 := Person{Name: "Charlie", Age: 35}
    
    result = list_tools.InList(p2, []Person{p1, p2, p3})
    fmt.Println("p2 是否在列表中:", result) // 输出: true
}
```

### 查找两个切片的差集

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/list_tools"
)

func main() {
    // 查找字符串切片的差集
    slice1 := []string{"apple", "banana", "orange", "grape"}
    slice2 := []string{"apple", "orange", "kiwi"}
    
    diff := list_tools.FindDifferenceNotInSlice(slice1, slice2)
    fmt.Println("slice1中不在slice2中的元素:", diff) // 输出: [banana grape]
    
    // 查找数字切片的差集
    nums1 := []int{1, 2, 3, 4, 5}
    nums2 := []int{1, 3, 5, 7, 9}
    
    diffNums := list_tools.FindDifferenceNotInSlice(nums1, nums2)
    fmt.Println("nums1中不在nums2中的元素:", diffNums) // 输出: [2 4]
}
```

### 切片去重

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/list_tools"
)

func main() {
    // 字符串切片去重
    strList := []string{"apple", "banana", "apple", "orange", "banana", "grape"}
    
    uniqueList := list_tools.UniqList(strList)
    fmt.Println("去重后的列表:", uniqueList) // 输出: [apple banana orange grape]
    
    // 数字切片去重
    numList := []int{1, 2, 3, 2, 4, 1, 5, 3}
    
    uniqueNums := list_tools.UniqList(numList)
    fmt.Println("去重后的数字列表:", uniqueNums) // 输出: [1 2 3 4 5]
}
```

### 切片分割

```go
package main

import (
    "fmt"
    "github.com/otkinlife/go_tools/list_tools"
)

func main() {
    // 将切片分割成多个子切片，每个子切片最多包含3个元素
    items := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
    
    chunks, err := list_tools.SplitList(items, 3)
    if err != nil {
        fmt.Println("分割失败:", err)
        return
    }
    
    fmt.Println("分割后的子切片:")
    for i, chunk := range chunks {
        fmt.Printf("第%d个子切片: %v\n", i+1, chunk)
    }
    // 输出:
    // 第1个子切片: [a b c]
    // 第2个子切片: [d e f]
    // 第3个子切片: [g h]
    
    // 分割数字切片
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    numChunks, err := list_tools.SplitList(numbers, 4)
    if err != nil {
        fmt.Println("分割失败:", err)
        return
    }
    
    fmt.Println("分割后的数字子切片:")
    for i, chunk := range numChunks {
        fmt.Printf("第%d个子切片: %v\n", i+1, chunk)
    }
    // 输出:
    // 第1个子切片: [1 2 3 4]
    // 第2个子切片: [5 6 7 8]
    // 第3个子切片: [9 10]
}
```

## 函数说明

### InList 函数

```go
func InList[T comparable](one T, list []T) bool
```

- `one`: 要查找的元素
- `list`: 要搜索的切片
- 返回值: 布尔值，表示元素是否在切片中

### FindDifferenceNotInSlice 函数

```go
func FindDifferenceNotInSlice[T comparable](slice1, slice2 []T) []T
```

- `slice1`: 第一个切片
- `slice2`: 第二个切片
- 返回值: 包含在`slice1`中但不在`slice2`中的元素的新切片

### UniqList 函数

```go
func UniqList[T comparable](list []T) []T
```

- `list`: 输入切片
- 返回值: 去重后的新切片

### SplitList 函数

```go
func SplitList[T any](list []T, n int) ([][]T, error)
```

- `list`: 要分割的切片
- `n`: 每个子切片的最大元素数量
- 返回值:
    - 分割后的子切片数组
    - 错误信息（如果有）

## 注意事项

- 所有函数都使用Go泛型实现，要求Go 1.18或更高版本
- `InList`和`FindDifferenceNotInSlice`函数要求元素类型满足`comparable`约束
- `SplitList`函数适用于任何类型的切片
- `SplitList`函数的参数`n`必须大于0，否则会返回错误