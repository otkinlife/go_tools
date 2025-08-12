package http_tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetHttpCode 获取 HTTP 状态码
// return: HTTP 状态码
func (r *ReqClient) GetHttpCode() int {
	return r.response.StatusCode
}

// GetResponseHeader 获取响应头
func (r *ReqClient) GetResponseHeader() http.Header {
	return r.response.Header
}

// GetBody 获取响应体
// return: 响应体
func (r *ReqClient) GetBody() ([]byte, error) {
	body, err := io.ReadAll(r.response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetBodyString 获取响应体字符串
// return: 响应体字符串
func (r *ReqClient) GetBodyString() string {
	body, err := r.GetBody()
	if err != nil {
		return fmt.Sprintf("Error reading body: %v", err)
	}
	if body == nil {
		return ""
	}
	return string(body)
}

// GetBodyReadCloser 获取响应体读取器
func (r *ReqClient) GetBodyReadCloser() (io.ReadCloser, error) {
	return r.response.Body, nil
}

// LoadBody 将body数据写入data
// data: 待写入数据对象
// return: 错误
func (r *ReqClient) LoadBody(data any) error {
	body, err := r.GetBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &data)
}
