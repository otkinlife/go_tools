package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/otkinlife/go_tools/img"
	"image"
	_ "image/jpeg" // 注册 JPEG 格式解码器
	_ "image/png"  // 注册 PNG 格式解码器
	"io"
	"net/http"
	"os"
	"strings"
)

type DImgRet struct {
	Dir      string // 图片下载的路径
	FileName string // 图片文件名
	Filepath string // 图片文件路径
	Err      error  // 错误信息
	Format   string // 图片格式
}

func DownloadImage(url, targetDir string) DImgRet {
	if targetDir == "" {
		targetDir = os.TempDir()
	}
	if !strings.HasSuffix(targetDir, "/") {
		targetDir = targetDir + "/"
	}
	ret := DImgRet{
		Dir:    targetDir,
		Err:    nil,
		Format: img.Unknown,
	}

	// 发送 HTTP 请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		ret.Err = fmt.Errorf("failed to download image: %w", err)
		return ret
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		ret.Err = fmt.Errorf("bad status: %s", resp.Status)
		return ret
	}

	// 使用 TeeReader 来同时读取响应体并写入文件
	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	// 尝试解码图片
	_, format, err := image.Decode(tee)
	if err != nil {
		ret.Err = fmt.Errorf("failed to decode image: %w", err)
		return ret
	}

	// 检查图片格式
	if !img.IsImage(format) {
		ret.Err = fmt.Errorf("invalid image format")
		return ret
	}
	ret.Format = format

	// 创建目标路径的所有目录
	err = os.MkdirAll(ret.Dir, os.ModePerm)
	if err != nil {
		ret.Err = fmt.Errorf("failed to create directories: %w", err)
		return ret
	}

	// 构建文件名
	hash := md5.New()
	if _, err := io.Copy(hash, &buf); err != nil {
		ret.Err = fmt.Errorf("failed to calculate MD5: %w", err)
		return ret
	}
	md5Value := fmt.Sprintf("%x", hash.Sum(nil))
	ret.FileName = fmt.Sprintf("%s.%s", md5Value, format)
	ret.Filepath = fmt.Sprintf("%s%s", ret.Dir, ret.FileName)
	// 创建文件
	out, err := os.Create(ret.Filepath)
	if err != nil {
		ret.Err = fmt.Errorf("failed to create file: %w", err)
		return ret
	}
	defer out.Close()

	return ret
}
