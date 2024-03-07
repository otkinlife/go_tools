package http

import (
	"github.com/moul/http2curl"
	"net/http"
)

// ConvertToCurlString 将请求转换为 cURL 命令
// r: 请求
// return: cURL 命令
func ConvertToCurlString(r *http.Request) (string, error) {
	curlCmd, err := http2curl.GetCurlCommand(r)
	if err != nil {
		return "", err
	}
	return curlCmd.String(), nil
}
