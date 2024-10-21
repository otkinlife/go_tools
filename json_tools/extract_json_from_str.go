package json_tools

import (
	"errors"
)

// ExtractJsonFromStr 从字符串中提取JSON子串
func ExtractJsonFromStr(input string) (string, error) {
	var jsonStart, jsonEnd int
	var found bool
	stack := 0

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '{', '[':
			if stack == 0 {
				jsonStart = i
				found = true
			}
			stack++
		case '}', ']':
			stack--
			if stack == 0 && found {
				jsonEnd = i
				return input[jsonStart : jsonEnd+1], nil
			}
		}
	}

	if !found {
		return "", errors.New("no JSON found in input string")
	}

	return "", errors.New("incomplete JSON string")
}
