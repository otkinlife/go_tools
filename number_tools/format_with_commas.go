package number_tools

import (
	"fmt"
	"math"
	"strings"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func FormatNumberWithCommas[T Number](originalNum T, split string, decimalPlaces int) string {
	// 将数字转换为浮点数
	number := float64(originalNum)

	// 根据指定的小数点位数进行四舍五入
	factor := math.Pow(10, float64(decimalPlaces))
	rounded := math.Round(number*factor) / factor

	// 将数字分为整数部分和小数部分
	integerPart := int(rounded)
	decimalPart := rounded - float64(integerPart)

	// 将整数部分转换为字符串，并添加千分位分隔符
	integerStr := fmt.Sprintf("%d", integerPart)
	var withCommas []string
	for i, j := len(integerStr), 0; i > 0; i, j = i-3, j+3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		withCommas = append([]string{integerStr[start:i]}, withCommas...)
	}
	formattedInteger := strings.Join(withCommas, split)

	// 如果指定的小数点位数为0，则只返回整数部分
	if decimalPlaces == 0 {
		return formattedInteger
	}

	// 将小数部分转换为字符串，确保它有指定的位数
	formattedDecimal := fmt.Sprintf("%.*f", decimalPlaces, decimalPart)
	formattedDecimal = strings.TrimPrefix(formattedDecimal, "0")

	// 将整数部分和小数部分组合成结果字符串
	result := fmt.Sprintf("%s%s", formattedInteger, formattedDecimal)
	return result
}
