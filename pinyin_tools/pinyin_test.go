package pinyin_tools

import "testing"

func TestConvertToPinyin(t *testing.T) {
	// 初始化词典（必须调用）
	err := InitDict()
	if err != nil {
		t.Fatal("初始化词典失败:", err)
	}

	ret := ConvertToPinyin("你好世s界", "")
	t.Log("ret:", ret)
	ret = ConvertToPinyin("你好世s界", "-")
	t.Log("ret:", ret)
	list := ConvertToPinyinList("你好世hello界")
	t.Log("list:", list)
}

func TestPolyphoneWords(t *testing.T) {
	// 初始化词典（必须调用）
	err := InitDict()
	if err != nil {
		t.Fatal("初始化词典失败:", err)
	}

	// 测试多音字
	testCases := []struct {
		input    string
		expected string // 期望的读音
		desc     string
	}{
		{"银行", "yin-hang", "银行(hang)"},
		{"行走", "xing-zou", "行走(xing)"},
		{"重庆", "chong-qing", "重庆(chong)"},
		{"重要", "zhong-yao", "重要(zhong)"},
		{"重复", "chong-fu", "重复(chong)"},
		{"乐观", "le-guan", "乐观(le)"},
		{"音乐", "yin-yue", "音乐(yue)"},
		{"长度", "chang-du", "长度(chang)"},
		{"成长", "cheng-zhang", "成长(zhang)"},
		{"都会区", "dou-hui-qu", "都会区(默认)"},
	}

	for _, tc := range testCases {
		result := ConvertToPinyin(tc.input, "-")
		t.Logf("%s: 输入='%s', 实际输出='%s', 期望='%s'",
			tc.desc, tc.input, result, tc.expected)
	}

	// 测试自定义词典
	t.Log("\n=== 测试自定义词典 ===")
	AddCustomDict("都会区", "du hui qu")
	result := ConvertToPinyin("都会区", "-")
	t.Logf("自定义词典后: 输入='%s', 输出='%s'", "都会区", result)
}

func TestCompareSimpleVsSmart(t *testing.T) {
	// 初始化词典
	err := InitDict()
	if err != nil {
		t.Fatal("初始化词典失败:", err)
	}

	testCases := []string{
		"银行", "行走", "重庆", "重要", "音乐", "乐观",
	}

	t.Log("\n=== 对比简单转换vs智能转换 ===")
	for _, text := range testCases {
		simple := ConvertToPinyinSimple(text, "-")
		smart := ConvertToPinyin(text, "-")
		t.Logf("'%s': 简单转换='%s', 智能转换='%s'", text, simple, smart)
	}
}
