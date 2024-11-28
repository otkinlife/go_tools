package http_tools

import (
	"log"
	"net/http"
)

type ReqClient struct {
	client      *http.Client
	req         *http.Request
	response    *http.Response
	isPrintCurl bool
}

func NewReqClient(method, url string) (*ReqClient, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return &ReqClient{
		client: &http.Client{},
		req:    req,
	}, nil
}

// Send 发送请求
// method: 请求方法
// urlStr: 请求地址
// return: 错误
func (r *ReqClient) Send() error {
	var err error
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
