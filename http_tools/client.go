package http_tools

import (
	"log"
	"net/http"
	"net/url"
)

type ReqClient struct {
	client      *http.Client
	req         *http.Request
	response    *http.Response
	isPrintCurl bool
}

func NewReqClient(method, url string) *ReqClient {
	req, _ := http.NewRequest(method, url, nil)
	return &ReqClient{
		client: &http.Client{},
		req:    req,
	}
}

// Get 发送 GET 请求
// urlStr: 请求地址
// return: 错误
func (r *ReqClient) Get(urlStr string) error {
	return r.Send("GET", urlStr)
}

// Post 发送 POST 请求
// urlStr: 请求地址
// return: 错误
func (r *ReqClient) Post(urlStr string) error {
	return r.Send("POST", urlStr)
}

// GetRetry 发送 GET 请求，重试
// urlStr: 请求地址
// tryCount: 重试次数
// return: 错误
func (r *ReqClient) GetRetry(urlStr string, tryCount int) error {
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
func (r *ReqClient) PostRetry(urlStr string, tryCount int) error {
	var err error
	for i := 0; i < tryCount; i++ {
		err = r.Post(urlStr)
		if err == nil && r.GetHttpCode() == 200 {
			break
		}
	}
	return err
}

// Send 发送请求
// method: 请求方法
// urlStr: 请求地址
// return: 错误
func (r *ReqClient) Send(urlStr string) error {
	var err error
	r.req.Method = method
	r.req.URL, err = url.Parse(urlStr)
	if err != nil {
		return err
	}
	if r.isPrintCurl {
		log.Println(ConvertToCurlString(*r.req))
	}
	resp, err := r.client.Do(r.req)
	if err != nil {
		return err
	}
	r.response = resp
	return nil
}

// Close 关闭请求
func (r *ReqClient) Close() {
	if r.response != nil {
		err := r.response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
