package http_tools

import "testing"

func TestSendGet(t *testing.T) {
	t.Log("TestSendGet")
	client := NewReqClient("GET", "http://www.baidu.com")
	err := client.Send("GET", "http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
}
