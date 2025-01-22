package img

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// AvgSplitImg 将矩形图片平均切割成 count 份，并返回切割后的图片路径列表
// count: 切割份数
// filePath: 图片文件路径
// targetDir: 切割后的图片保存目录
// return: 切割后的图片路径列表
// return: 错误
func AvgSplitImg(count int, filePath, targetDir string) ([]string, error) {
	// 打开图片文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	// 解码图片
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// 获取图片的尺寸
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// 确保 count 是完全平方数
	sqrtCount := int(math.Sqrt(float64(count)))
	if sqrtCount*sqrtCount != count {
		return nil, fmt.Errorf("count must be a perfect square")
	}

	// 计算每块的大小
	partWidth := width / sqrtCount
	partHeight := height / sqrtCount

	// 创建输出目录
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	// 分割图片并保存
	var paths []string
	for i := 0; i < sqrtCount; i++ {
		for j := 0; j < sqrtCount; j++ {
			// 计算每块的边界
			rect := image.Rect(i*partWidth, j*partHeight, (i+1)*partWidth, (j+1)*partHeight)
			part := image.NewRGBA(rect.Bounds())
			draw.Draw(part, part.Bounds(), img, rect.Min, draw.Src)

			// 创建输出文件名
			baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
			outFileName := fmt.Sprintf("%s_part_%d_%d.png", baseName, i, j)
			outFilePath := filepath.Join(targetDir, outFileName)

			// 创建输出文件
			outFile, err := os.Create(outFilePath)
			if err != nil {
				return nil, fmt.Errorf("failed to create output file: %v", err)
			}
			defer outFile.Close()

			// 保存分割后的图片
			if format == "jpeg" || format == "jpg" {
				err = jpeg.Encode(outFile, part, nil)
			} else {
				err = png.Encode(outFile, part)
			}
			if err != nil {
				return nil, fmt.Errorf("failed to encode image: %v", err)
			}

			paths = append(paths, outFilePath)
		}
	}

	return paths, nil
}

// AvgSplitImgFromUrl 从 URL 下载图片，切割并保存到指定目录
func AvgSplitImgFromUrl(count int, url, targetDir string) ([]string, error) {
	// 从 URL 下载图片
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "downloaded_image.*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// 将下载的内容写入临时文件
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to save downloaded image: %v", err)
	}

	// 关闭临时文件
	if err := tmpFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close temp file: %v", err)
	}

	// 调用 AvgSplitImg 方法进行图片切割
	return AvgSplitImg(count, tmpFile.Name(), targetDir)
}
