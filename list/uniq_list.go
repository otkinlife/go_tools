package list

func UniqList[T comparable](list []T) []T {
	newList := make([]T, 0)
	uniqueMap := make(map[T]bool)
	for _, item := range list {
		if _, ok := uniqueMap[item]; !ok {
			newList = append(newList, item)
			uniqueMap[item] = true
		}
	}
	return newList
}
