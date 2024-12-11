package http_tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// DownloadFile 下载文件
// filePath: 要保存的文件路径（包括文件名）
// chunkCount: 分块下载的块数: 最大为 32, 0 表示不分块下载
// return: 错误
func (r *ReqClient) DownloadFile(filePath string, chunkCount int) error {
	if chunkCount > 32 {
		chunkCount = 32
	}

	err := r.Send()
	if err != nil {
		return err
	}

	if r.GetHttpCode() != http.StatusOK {
		return fmt.Errorf("server returned %d status", r.GetHttpCode())
	}

	size, err := strconv.Atoi(r.response.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if chunkCount > 1 && size > chunkCount {
		return r.downloadFileChunks(file, size, chunkCount)
	}

	_, err = io.Copy(file, r.response.Body)
	return err
}

func (r *ReqClient) downloadFileChunks(file *os.File, size, chunkCount int) error {
	chunkSize := size / chunkCount
	var wg sync.WaitGroup
	wg.Add(chunkCount)

	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == chunkCount-1 {
			end = size
		}

		go func(start, end int) {
			defer wg.Done()

			r.req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end-1))
			err := r.Send()
			if err != nil {
				fmt.Println(err)
				return
			}

			_, _ = file.Seek(int64(start), 0)
			_, _ = io.Copy(file, r.response.Body)
		}(start, end)
	}

	wg.Wait()
	return nil
}
