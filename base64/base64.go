package base64

import "encoding/base64"

// Encode Base64加密
// s: 待加密字符串
// return: 加密后的字符串
func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Decode Base64解密
// s: 待解密字符串
// return: 解密后的字符串
func Decode(s string) (string, error) {
	b, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		return "", e
	}
	return string(b), nil
}
