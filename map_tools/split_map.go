package list

import "fmt"

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
