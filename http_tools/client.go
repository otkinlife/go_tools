package http_tools

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ReqClient struct {
	client      *http.Client
	req         *http.Request
	response    *http.Response
	isPrintCurl bool
	formFields  map[string]string
	files       []fileField
}

type fileField struct {
	fieldName string
	filePath  string
}

func NewReqClient(method, url string) (*ReqClient, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return &ReqClient{
		client:     &http.Client{},
		req:        req,
		formFields: make(map[string]string),
		files:      make([]fileField, 0),
	}, nil
}

// Send 发送请求
// method: 请求方法
// urlStr: 请求地址
// return: 错误
func (r *ReqClient) Send() error {
	// 构建multipart请求体
	if err := r.buildMultipartRequest(); err != nil {
		return err
	}

	var err error
	if r.isPrintCurl {
		// 打印curl命令，使用文件路径格式
		curlCmd, err := r.ConvertToCurlWithFiles()
		if err != nil {
			log.Printf("Failed to generate curl command: %v", err)
		} else {
			log.Println(curlCmd)
		}
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

// buildMultipartRequest 构建multipart请求体
func (r *ReqClient) buildMultipartRequest() error {
	if len(r.formFields) == 0 && len(r.files) == 0 {
		return nil
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加表单字段
	for key, value := range r.formFields {
		err := writer.WriteField(key, value)
		if err != nil {
			return err
		}
	}

	// 添加文件
	for _, fileField := range r.files {
		file, err := os.Open(fileField.filePath)
		if err != nil {
			return err
		}

		part, err := writer.CreateFormFile(fileField.fieldName, filepath.Base(fileField.filePath))
		if err != nil {
			_ = file.Close()
			return err
		}

		_, err = io.Copy(part, file)
		_ = file.Close()
		if err != nil {
			return err
		}
	}

	err := writer.Close()
	if err != nil {
		return err
	}

	r.req.Body = io.NopCloser(body)
	r.req.Header.Set("Content-Type", writer.FormDataContentType())
	return nil
}

// ConvertToCurlWithFiles 将请求转换为curl命令，对文件使用@filepath格式
func (r *ReqClient) ConvertToCurlWithFiles() (string, error) {
	var parts []string

	// 基础curl命令
	parts = append(parts, "curl")

	// 请求方法
	if r.req.Method != "GET" {
		parts = append(parts, "-X", r.req.Method)
	}

	// URL
	parts = append(parts, fmt.Sprintf("'%s'", r.req.URL.String()))

	// 请求头
	for key, values := range r.req.Header {
		// 跳过Content-Type，我们会手动设置
		if key == "Content-Type" {
			continue
		}
		for _, value := range values {
			parts = append(parts, "-H", fmt.Sprintf("'%s: %s'", key, value))
		}
	}

	// 处理表单数据和文件
	if len(r.formFields) > 0 || len(r.files) > 0 {
		// 添加表单字段
		for key, value := range r.formFields {
			parts = append(parts, "-F", fmt.Sprintf("'%s=%s'", key, value))
		}

		// 添加文件（使用@filepath格式）
		for _, file := range r.files {
			parts = append(parts, "-F", fmt.Sprintf("'%s=@%s'", file.fieldName, file.filePath))
		}
	} else if r.req.Body != nil {
		// 如果有其他类型的请求体，使用--data-raw
		bodyBytes, err := io.ReadAll(r.req.Body)
		if err != nil {
			return "", err
		}
		// 重置请求体
		r.req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if len(bodyBytes) > 0 {
			parts = append(parts, "--data-raw", fmt.Sprintf("'%s'", string(bodyBytes)))
		}
	}

	return strings.Join(parts, " "), nil
}
