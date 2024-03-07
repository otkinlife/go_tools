package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type Req struct {
	client   *http.Client
	req      *http.Request
	response *http.Response
}

func NewReq() *Req {
	return &Req{
		client: &http.Client{},
		req:    &http.Request{},
	}
}

// SetHeaders 设置请求头
// headers: 请求头
func (r *Req) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		r.req.Header.Set(key, value)
	}
}

// SetQuery 设置请求参数
// query: 请求参数
func (r *Req) SetQuery(query map[string]string) {
	q := url.Values{}
	for key, value := range query {
		q.Add(key, value)
	}
	r.req.URL.RawQuery = q.Encode()
}

// SetForm 设置表单
// form: 表单
func (r *Req) SetForm(form map[string]string) {
	f := url.Values{}
	for key, value := range form {
		f.Add(key, value)
	}
	r.req.PostForm = f
	r.req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

// SetJson 设置 JSON 请求体
// jsonObject: JSON 对象
func (r *Req) SetJson(jsonObject any) error {
	jsonValue, err := json.Marshal(jsonObject)
	if err != nil {
		return err
	}
	r.req.Body = io.NopCloser(bytes.NewReader(jsonValue))
	r.req.Header.Set("Content-Type", "application/json")
	return nil
}

// SetTimeout 设置超时时间
// timeout: 超时时间
func (r *Req) SetTimeout(timeout time.Duration) {
	r.client.Timeout = timeout
}

// Get 发送 GET 请求
// urlStr: 请求地址
// return: 错误
func (r *Req) Get(urlStr string) error {
	return r.Send("GET", urlStr)
}

// Post 发送 POST 请求
// urlStr: 请求地址
// return: 错误
func (r *Req) Post(urlStr string) error {
	return r.Send("POST", urlStr)
}

// GetRetry 发送 GET 请求，重试
// urlStr: 请求地址
// tryCount: 重试次数
// return: 错误
func (r *Req) GetRetry(urlStr string, tryCount int) error {
	var err error
	for i := 0; i < tryCount; i++ {
		err = r.Get(urlStr)
		if err == nil && r.GetHttpCode() == 200 {
			break
		}
	}
	return err
}

// PostRetry 发送 POST 请求，重试
// urlStr: 请求地址
// tryCount: 重试次数
func (r *Req) PostRetry(urlStr string, tryCount int) error {
	var err error
	for i := 0; i < tryCount; i++ {
		err = r.Post(urlStr)
		if err == nil && r.GetHttpCode() == 200 {
			break
		}
	}
	return err
}

// UploadFile 上传文件
// fieldName: 字段名
// filePath: 文件路径
// return: 错误
func (r *Req) UploadFile(fieldName, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, filepath.Base(filePath))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return err
	}

	r.req.Body = io.NopCloser(body)
	r.req.Header.Set("Content-Type", writer.FormDataContentType())
	return nil
}

// DownloadFile 下载文件
// url: 请求地址
// filePath: 要保存的文件路径（包括文件名）
// chunkCount: 分块下载的块数: 最大为 32, 0 表示不分块下载
// return: 错误
func (r *Req) DownloadFile(url, filePath string, chunkCount int) error {
	if chunkCount > 32 {
		chunkCount = 32
	}

	err := r.Get(url)
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
		return r.downloadFileChunks(url, file, size, chunkCount)
	}

	_, err = io.Copy(file, r.response.Body)
	return err
}

func (r *Req) downloadFileChunks(url string, file *os.File, size, chunkCount int) error {
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
			err := r.Get(url)
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

// Send 发送请求
// method: 请求方法
// urlStr: 请求地址
// return: 错误
func (r *Req) Send(method, urlStr string) error {
	var err error
	r.req.Method = method
	r.req.URL, err = url.Parse(urlStr)
	if err != nil {
		return err
	}

	resp, err := r.client.Do(r.req)
	if err != nil {
		return err
	}
	r.response = resp
	return nil
}

// GetHttpCode 获取 HTTP 状态码
// return: HTTP 状态码
func (r *Req) GetHttpCode() int {
	return r.response.StatusCode
}

// GetBody 获取响应体
// return: 响应体
func (r *Req) GetBody() ([]byte, error) {
	body, err := io.ReadAll(r.response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// LoadBody 加载响应体
// data: 数据
// return: 错误
func (r *Req) LoadBody(data any) error {
	body, err := r.GetBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &data)
}

// Close 关闭请求
func (r *Req) Close() {
	err := r.response.Body.Close()
	if err != nil {
		log.Println(err)
	}
}

// GetCurlString 获取 cURL 命令
func (r *Req) GetCurlString() (string, error) {
	return ConvertToCurlString(r.req)
}
