# map_tools

`map_tools` 是一个用于将 Go 语言中的 `map` 数据结构分割为数组或多个较小的 `map` 的工具包。该工具包包含两个主要函数：`SplitMap2List` 和 `SplitMap2Map`。

## 函数说明

### SplitMap2List

将一个 `map` 分割为多个数组。

#### 函数签名

```go
func SplitMap2List[T any, K comparable](m map[K]T, chunkSize int) ([][]T, error)
```

#### 参数

- `m`: 要分割的 `map`。
- `chunkSize`: 每个数组的长度。

#### 返回值

- `[][]T`: 分割后的数组切片。
- `error`: 错误信息，如果 `chunkSize` 小于等于 0 或者 `map` 为空。

#### 示例

```go
package main

import (
	"fmt"
	"map_tools"
)

func main() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	chunkSize := 2
	result, err := map_tools.SplitMap2List(m, chunkSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result) // Output: [["a", "b"], ["c", "d"]]
}
```

### SplitMap2Map

将一个 `map` 分割为多个较小的 `map`。

#### 函数签名

```go
func SplitMap2Map[T any, K comparable](m map[K]T, chunkSize int) ([]map[K]T, error)
```

#### 参数

- `m`: 要分割的 `map`。
- `chunkSize`: 每个 `map` 的长度。

#### 返回值

- `[]map[K]T`: 分割后的 `map` 切片。
- `error`: 错误信息，如果 `chunkSize` 小于等于 0。

#### 示例

```go
package main

import (
	"fmt"
	"map_tools"
)

func main() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	chunkSize := 2
	result, err := map_tools.SplitMap2Map(m, chunkSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, chunk := range result {
		fmt.Println(chunk)
	}
	// Output: 
	// map[1:a 2:b]
	// map[3:c 4:d]
}
```

## 安装

将 `map_tools` 包下载到你的项目中：

```shell
go get -u github.com/yourusername/map_tools
```

## 使用
```go
m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
		6: "f",
		7: "g",
	}
chunkSize := 2
result, err := SplitMap2List(m, chunkSize)
result, err := TestSplitMap2Map(m, chunkSize)
```