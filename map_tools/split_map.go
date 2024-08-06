package map_tools

import (
	"errors"
	"fmt"
)

// SplitMap 将map分割为数组
// list: 要分割的map
// chunkSize: 每个数组的长度
func SplitMap[T any, K comparable](list map[K]T, chunkSize int) ([][]T, error) {
	if chunkSize <= 0 {
		return nil, fmt.Errorf("chunkSize must be greater than 0")
	}

	// Convert map to slice of values
	values := make([]T, 0, len(list))
	for _, value := range list {
		values = append(values, value)
	}

	total := len(values)
	if total == 0 {
		return nil, fmt.Errorf("map is empty")
	}

	// Split values into chunks
	var ret [][]T
	for i := 0; i < total; i += chunkSize {
		end := i + chunkSize
		if end > total {
			end = total
		}
		ret = append(ret, values[i:end])
	}

	return ret, nil
}

// SplitMap2Map 将map分割为map
// m: 要分割的map
// chunkSize: 每个map的长度
func SplitMap2Map[T any, K comparable](m map[K]T, chunkSize int) ([]map[K]T, error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunkSize must be greater than 0")
	}

	if len(m) == 0 {
		return []map[K]T{}, nil
	}

	var result []map[K]T
	currentChunk := make(map[K]T)
	count := 0

	for k, v := range m {
		currentChunk[k] = v
		count++
		if count == chunkSize {
			result = append(result, currentChunk)
			currentChunk = make(map[K]T)
			count = 0
		}
	}

	if count > 0 {
		result = append(result, currentChunk)
	}

	return result, nil
}
