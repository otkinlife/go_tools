package http_tools

import "testing"

func TestSendGet(t *testing.T) {
	t.Log("TestSendGet")
	client, _ := NewReqClient("GET", "http://www.baidu.com")
	err := client.Send()
	if err != nil {
		t.Error(err)
	}
}

func TestDownloadFile(t *testing.T) {
	t.Log("TestDownloadFile")
	client, _ := NewReqClient("GET", "https://bi-base-data.oss-accelerate-overseas.aliyuncs.com/maxcompute/billing/appstore_financial/app%3D102/month%3D2024-06/data_C4Q66Q773B_0.csv?OSSAccessKeyId=LTAI5tQSpxrfJtR1KDsNZwxh&Expires=1732777491&Signature=3mq%2B%2BmbWUOU7wZ76aed5RMKYlXk%3D")
	err := client.DownloadFile("./baidu.html", 0)
	if err != nil {
		t.Error(err)
	}
}
