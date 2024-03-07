package list

// InList 判断一个字符串是否在一个字符串数组中
// one: 待判断字符串
// list: 字符串数组
// return: 是否在数组中
func InList[T comparable](one T, list []T) bool {
	for _, item := range list {
		if item == one {
			return true
		}
	}
	return false
}
