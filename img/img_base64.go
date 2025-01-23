package img

import (
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
)

func LocalImageToDataURL(imagePath string) (string, error) {
	// Guess the MIME type of the image based on the file extension
	ext := filepath.Ext(imagePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // Default MIME type if none is found
	}

	// Read and encode the image file
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	base64EncodedData := base64.StdEncoding.EncodeToString(imageData)

	// Construct the data URL
	dataURL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64EncodedData)
	return dataURL, nil
}

func DecodeBase642Img(base64Str, imgPath string) error {
	// 解码 Base64 字符串
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	// 获取文件的目录路径
	dir := filepath.Dir(imgPath)

	// 创建目录（如果不存在）
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// 打开或创建文件
	file, err := os.OpenFile(imgPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将图像数据写入文件
	_, err = file.Write(imageData)
	if err != nil {
		return err
	}

	return nil
}
