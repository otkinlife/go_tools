package pinyin_tools

import (
	"strings"

	"github.com/go-ego/gpy/phrase"
	"github.com/mozillazg/go-pinyin"
)

// ConvertToPinyin 将汉字转换为拼音（支持多音字处理）
// str: 待转换汉字
// split: 拼音之间的分隔符
// return: 转换后的拼音
func ConvertToPinyin(str, split string) string {
	// 使用gpy库进行智能转换，支持多音字处理
	result := phrase.Paragraph(str)

	// 如果split不是空格，需要替换分隔符
	if split != " " && split != "" {
		result = strings.ReplaceAll(result, " ", split)
	} else if split == "" {
		// 如果不需要分隔符，去掉所有空格
		result = strings.ReplaceAll(result, " ", "")
	}

	return result
}

// ConvertToPinyinList 将汉字转换为拼音数组
// str: 待转换汉字
// return: 转换后的拼音数组
func ConvertToPinyinList(str string) []string {
	// 使用gpy库进行智能转换
	result := phrase.Paragraph(str)
	return strings.Fields(result) // 按空格分割成数组
}

// ConvertToPinyinSimple 使用原始go-pinyin库进行简单转换（逐字转换）
// str: 待转换汉字
// split: 拼音之间的分隔符
// return: 转换后的拼音
func ConvertToPinyinSimple(str, split string) string {
	newStr := strings.Join(ConvertToPinyinListSimple(str), split)
	return newStr
}

// ConvertToPinyinListSimple 使用原始go-pinyin库进行简单转换（逐字转换）
// str: 待转换汉字
// return: 转换后的拼音数组
func ConvertToPinyinListSimple(str string) []string {
	a := pinyin.NewArgs()
	pinyinStr := pinyin.Pinyin(str, a)

	// 将拼音数组转换为字符串
	newList := make([]string, 0)
	idx := 0
	for _, c := range str {
		if isChinese(c) {
			if idx < len(pinyinStr) && len(pinyinStr[idx]) > 0 {
				newList = append(newList, pinyinStr[idx][0])
			} else {
				newList = append(newList, string(c))
			}
			idx++
		} else {
			newList = append(newList, string(c))
		}
	}
	return newList
}

// AddCustomDict 添加自定义拼音词典
// word: 词语
// pinyin: 对应的拼音（用空格分隔）
func AddCustomDict(word, pinyin string) {
	phrase.AddDict(word, pinyin)
}

// InitDict 初始化内置词典（必须调用才能启用多音字处理）
func InitDict() error {
	return phrase.LoadGseDictEmbed("zh")
}

// isChinese 判断字符是否为汉字
// c: 待判断字符
// return: 是否为汉字
// 汉字的Unicode编码范围为0x4E00~0x9FA5
func isChinese(c rune) bool {
	return c >= 0x4E00 && c <= 0x9FA5
}
