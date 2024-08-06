package list

// FindDifferenceNotInSlice2 返回slice1中不在slice2中的元素
func FindDifferenceNotInSlice2[T comparable](slice1, slice2 []T) []T {
	elementMap := make(map[T]struct{})
	for _, elem := range slice2 {
		elementMap[elem] = struct{}{}
	}

	var result []T
	for _, elem := range slice1 {
		if _, found := elementMap[elem]; !found {
			result = append(result, elem)
		}
	}

	return result
}
