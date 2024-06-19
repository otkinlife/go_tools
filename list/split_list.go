package list

import "fmt"

func SplitList[T any](list []T, n int) ([][]T, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be greater than 0")
	}
	var ret [][]T
	for i := 0; i < len(list); i += n {
		end := i + n
		if end > len(list) {
			end = len(list)
		}
		ret = append(ret, list[i:end])
	}
	return ret, nil
}
