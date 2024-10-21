package json_tools

import "testing"

func TestExtractJsonFromStr(t *testing.T) {
	str := "根据给定的知识库内容和用户反馈,由于知识库为空,无法判断是否需要查询用户数据,因此输出如下:\n\n{\"intent\":\"query\",\"slots\":{\"query\":\"用户数据\"},\"session\":\"\"}"
	json, err := ExtractJsonFromStr(str)
	if err != nil {
		t.Errorf("ExtractJsonFromStr() error = %v", err)
		return
	}
	t.Logf("ExtractJsonFromStr() json = %v", json)
}
