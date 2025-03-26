package md5

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/otkinlife/go_tools/downloader"
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

// FileUrlGenerate 基于文件URL生成MD5
// url: 文件URL
func FileUrlGenerate(url, saveFileName string) (string, error) {
	err := downloader.DownloadFile(url, saveFileName, 8)
	if err != nil {
		return "", err
	}
	defer os.Remove(saveFileName)
	file, err := os.Open(saveFileName)
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
