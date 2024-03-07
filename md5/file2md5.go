package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// FileGenerate 基于文件生成MD5
// filePath: 文件路径
// return: MD5字符串
func FileGenerate(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
