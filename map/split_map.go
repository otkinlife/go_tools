package list

import "fmt"

func SplitMap[T any, K comparable](list map[K]T, n int) ([][]T, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be greater than 0")
	}

	// Convert map to slice of values
	values := make([]T, 0, len(list))
	for _, value := range list {
		values = append(values, value)
	}

	// Calculate the size of each chunk
	total := len(values)
	if total == 0 {
		return nil, fmt.Errorf("map is empty")
	}

	chunkSize := (total + n - 1) / n // This ensures we handle rounding up correctly

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
