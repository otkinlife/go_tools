package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// StringGenerate 基于字符串生成MD5
// input: 输入字符串
// return: MD5字符串
func StringGenerate(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
