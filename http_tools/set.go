package http_tools

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SetHeaders 设置请求头
// headers: 请求头
func (r *ReqClient) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		r.req.Header.Set(key, value)
	}
}

// SetQuery 设置请求参数
// query: 请求参数
func (r *ReqClient) SetQuery(query map[string]string) {
	q := url.Values{}
	for key, value := range query {
		q.Add(key, value)
	}
	r.req.URL.RawQuery = q.Encode()
}

// SetForm 设置表单
// form: 表单
func (r *ReqClient) SetForm(form map[string]string) {
	f := url.Values{}
	for key, value := range form {
		f.Add(key, value)
	}
	formData := f.Encode()
	r.req.Body = io.NopCloser(strings.NewReader(formData))
	r.req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}

// SetFile 上传文件
// fieldName: 字段名
// filePath: 文件路径
// return: 错误
func (r *ReqClient) SetFile(fieldName, filePath string) error {
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

// SetJson 设置 JSON 请求体
// jsonObject: JSON 对象
func (r *ReqClient) SetJson(jsonObject any) error {
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
func (r *ReqClient) SetTimeout(timeout time.Duration) {
	r.client.Timeout = timeout
}

// SetIsPrintCurl 设置是否打印 curl 命令
// isPrintCurl: 是否打印 curl 命令
func (r *ReqClient) SetIsPrintCurl(isPrintCurl bool) {
	r.isPrintCurl = isPrintCurl
}
