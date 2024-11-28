package http_tools

import (
	"encoding/json"
	"io"
)

// GetHttpCode 获取 HTTP 状态码
// return: HTTP 状态码
func (r *ReqClient) GetHttpCode() int {
	return r.response.StatusCode
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

// LoadBody 加载响应体
// data: 数据
// return: 错误
func (r *ReqClient) LoadBody(data any) error {
	body, err := r.GetBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &data)
}
