package downloader

import (
	"fmt"
	"github.com/otkinlife/go_tools/http_tools"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DownloadFile 下载文件
// url: 文件 URL
// filePath: 要保存的文件路径(如果带有文件名则保存为指定文件名, 否则保存为 URL 中的文件名)
// chunkCount: 分块下载的块数: 最大为 32, 0 表示不分块下载
// return: 错误
func DownloadFile(url string, filePath string, chunkCount int) error {
	// 参数验证
	if chunkCount < 0 || chunkCount > 32 {
		return fmt.Errorf("invalid chunkCount: %d, must be between 0 and 32", chunkCount)
	}

	// 处理文件路径
	dirPath, fileName, err := processFilePath(url, filePath)
	if err != nil {
		return err
	}

	// 确保目录存在
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}

	// 完整文件路径
	fullPath := filepath.Join(dirPath, fileName)

	// 创建临时文件，下载完成后再重命名
	tempPath := fullPath + ".download"
	file, err := os.Create(tempPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		file.Close()
		// 如果函数返回错误，删除临时文件
		if err != nil {
			os.Remove(tempPath)
		}
	}()

	// 获取文件大小和检查是否支持Range请求
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Head(url)
	if err != nil {
		return fmt.Errorf("failed to send HEAD request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-OK status: %s", resp.Status)
	}

	// 检查Content-Length
	contentLengthStr := resp.Header.Get("Content-Length")
	if contentLengthStr == "" {
		// 如果服务器没有提供Content-Length，则使用非分块下载
		chunkCount = 0
	}

	var fileSize int
	if contentLengthStr != "" {
		fileSize, err = strconv.Atoi(contentLengthStr)
		if err != nil {
			return fmt.Errorf("invalid Content-Length: %w", err)
		}
	}

	// 检查是否支持Range请求
	acceptRanges := resp.Header.Get("Accept-Ranges")
	if chunkCount > 1 && acceptRanges != "bytes" {
		// 服务器不支持Range请求，回退到非分块下载
		chunkCount = 0
	}

	if chunkCount <= 1 {
		// 不分块下载
		if err := downloadSingleChunk(url, file); err != nil {
			return err
		}
		// 关闭文件并重命名
		if err := file.Close(); err != nil {
			return fmt.Errorf("failed to close file: %w", err)
		}

		if err := os.Rename(tempPath, fullPath); err != nil {
			return fmt.Errorf("failed to rename temp file: %w", err)
		}

		return nil
	}
	// 分块下载
	return downloadMultipleChunks(url, file, fileSize, chunkCount, tempPath, fullPath)
}

// 处理文件路径，返回目录路径和文件名
func processFilePath(urlStr, filePath string) (string, string, error) {
	// 解析URL以提取正确的文件名
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", "", fmt.Errorf("invalid URL: %w", err)
	}

	// 从URL路径中提取文件名，忽略查询参数
	urlFileName := path.Base(parsedURL.Path)
	if urlFileName == "." || urlFileName == "/" || urlFileName == "" {
		return "", "", fmt.Errorf("cannot extract filename from URL: %s", urlStr)
	}

	// 处理URL编码的文件名
	urlFileName, err = url.QueryUnescape(urlFileName)
	if err != nil {
		return "", "", fmt.Errorf("failed to decode URL filename: %w", err)
	}

	// 检查filePath是否以路径分隔符结尾
	if strings.HasSuffix(filePath, "/") || strings.HasSuffix(filePath, "\\") {
		// filePath是目录
		return filePath, urlFileName, nil
	}

	// 检查filePath是否包含路径分隔符
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)

	if dir == "." && base == filePath {
		// filePath只是文件名，使用当前目录
		return ".", filePath, nil
	}

	// filePath包含目录和文件名
	return dir, base, nil
}

// 单块下载
func downloadSingleChunk(url string, file *os.File) error {
	reqClient, err := http_tools.NewReqClient("GET", url)
	if err != nil {
		return fmt.Errorf("failed to create request client: %w", err)
	}
	defer reqClient.Close()

	// 设置超时
	reqClient.SetTimeout(30 * time.Second)

	if err := reqClient.Send(); err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	body, err := reqClient.GetBody()
	if err != nil {
		return fmt.Errorf("failed to get response body: %w", err)
	}

	if _, err = file.Write(body); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// 多块下载
func downloadMultipleChunks(url string, file *os.File, fileSize, chunkCount int, tempPath, fullPath string) error {
	var wg sync.WaitGroup
	errChan := make(chan error, chunkCount) // 每个goroutine可能产生一个错误
	chunkSize := fileSize / chunkCount

	// 预分配文件大小，避免并发写入时的文件增长问题
	if err := file.Truncate(int64(fileSize)); err != nil {
		return fmt.Errorf("failed to allocate file size: %w", err)
	}

	for i := 0; i < chunkCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			start := i * chunkSize
			end := start + chunkSize - 1
			if i == chunkCount-1 {
				end = fileSize - 1
			}

			if err := downloadChunk(url, file, start, end); err != nil {
				select {
				case errChan <- fmt.Errorf("chunk %d download failed: %w", i, err):
				default:
				}
			}
		}(i)
	}

	// 等待所有下载完成
	doneChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	// 等待完成或错误
	select {
	case err := <-errChan:
		return err
	case <-doneChan:
		// 所有块下载完成，重命名文件
		if err := file.Close(); err != nil {
			return fmt.Errorf("failed to close file: %w", err)
		}
		if err := os.Rename(tempPath, fullPath); err != nil {
			return fmt.Errorf("failed to rename temp file: %w", err)
		}
		return nil
	}
}

// 下载单个块
func downloadChunk(url string, file *os.File, start, end int) error {
	reqClient, err := http_tools.NewReqClient("GET", url)
	if err != nil {
		return err
	}
	defer reqClient.Close()

	// 设置超时
	reqClient.SetTimeout(30 * time.Second)

	reqClient.SetHeaders(map[string]string{
		"Range": fmt.Sprintf("bytes=%d-%d", start, end),
	})

	if err := reqClient.Send(); err != nil {
		return err
	}

	body, err := reqClient.GetBody()
	if err != nil {
		return err
	}

	// 使用文件锁确保并发安全
	fileLock := &sync.Mutex{}
	fileLock.Lock()
	defer fileLock.Unlock()

	if _, err := file.Seek(int64(start), io.SeekStart); err != nil {
		return err
	}

	if _, err := file.Write(body); err != nil {
		return err
	}

	return nil
}
