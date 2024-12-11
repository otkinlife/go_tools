package downloader

import (
	"bytes"
	"fmt"
	"github.com/otkinlife/go_tools/http_tools"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// DownloadFile 下载文件
// url: 文件 URL
// filePath: 要保存的文件路径（包括文件名）
// chunkCount: 分块下载的块数: 最大为 32, 0 表示不分块下载
// return: 错误
func DownloadFile(url, filePath string, chunkCount int) error {
	if chunkCount < 0 || chunkCount > 32 {
		return fmt.Errorf("invalid chunkCount: %d, must be between 0 and 32", chunkCount)
	}

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取文件大小
	client := &http.Client{}
	resp, err := client.Head(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get file info: %s", resp.Status)
	}
	fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	if chunkCount == 0 {
		// 不分块下载
		reqClient, err := http_tools.NewReqClient("GET", url)
		if err != nil {
			return err
		}
		defer reqClient.Close()

		if err := reqClient.Send(); err != nil {
			return err
		}
		defer reqClient.Close()
		body, _ := reqClient.GetBody()
		_, err = io.Copy(file, bytes.NewReader(body))
		return err
	}

	// 分块下载
	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, 1)
	chunkSize := fileSize / chunkCount

	for i := 0; i < chunkCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			start := i * chunkSize
			end := start + chunkSize - 1
			if i == chunkCount-1 {
				end = fileSize - 1
			}

			reqClient, err := http_tools.NewReqClient("GET", url)
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}
			defer reqClient.Close()

			reqClient.SetHeaders(map[string]string{
				"Range": fmt.Sprintf("bytes=%d-%d", start, end),
			})
			if err := reqClient.Send(); err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}
			defer reqClient.Close()

			mu.Lock()
			file.Seek(int64(start), 0)
			body, _ := reqClient.GetBody()
			_, err = io.Copy(file, bytes.NewReader(body))
			mu.Unlock()

			if err != nil {
				select {
				case errCh <- err:
				default:
				}
			}
		}(i)
	}

	// 等待所有分块下载完成或出现错误
	go func() {
		wg.Wait()
		close(errCh)
	}()

	if err := <-errCh; err != nil {
		return err
	}

	return nil
}
