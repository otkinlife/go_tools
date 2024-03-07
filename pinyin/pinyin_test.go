package pinyin

import "testing"

func TestConvertToPinyin(t *testing.T) {
	ret := ConvertToPinyin("你好世s界", "")
	t.Log("ret:", ret)
	ret = ConvertToPinyin("你好世s界", "-")
	t.Log("ret:", ret)
}
