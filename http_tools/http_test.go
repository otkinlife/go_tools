package http_tools

import (
	"testing"
)

func TestSendGet(t *testing.T) {
	t.Log("TestSendGet")
	client, _ := NewReqClient("GET", "https://mpanso.me/DEMO.json")
	err := client.Send()
	if err != nil {
		t.Error(err)
	}
	t.Log(client.GetBody())
}
