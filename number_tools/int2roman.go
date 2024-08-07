package number_tools

import "strings"

// IntToRoman 将整数转换为罗马数字（1-3999）
// num: 整数
// return: 罗马数字
func IntToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	s := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman strings.Builder

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman.WriteString(s[i])
		}
	}

	return roman.String()
}
