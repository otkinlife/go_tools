package pinyin

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

// ConvertToPinyin 将汉字转换为拼音
// str: 待转换汉字
// split: 拼音之间的分隔符
// return: 转换后的拼音
func ConvertToPinyin(str, split string) string {
	a := pinyin.NewArgs()
	pinyinStr := pinyin.Pinyin(str, a)
	// 将拼音数组转换为字符串
	newList := make([]string, 0)
	idx := 0
	for _, c := range str {
		if isChinese(c) {
			newList = append(newList, pinyinStr[idx][0])
			idx++
		} else {
			newList = append(newList, string(c))
		}
	}
	newStr := strings.Join(newList, split)
	return newStr
}

// isChinese 判断字符是否为汉字
// c: 待判断字符
// return: 是否为汉字
// 汉字的Unicode编码范围为0x4E00~0x9FA5
func isChinese(c rune) bool {
	return c >= 0x4E00 && c <= 0x9FA5
}
