package list

import "testing"

type TItem struct {
	A int
	B string
}

func TestInList(t *testing.T) {
	// 字符串
	ret := InList("a", []string{"a", "b", "c"})
	t.Log("ret:", ret)

	// 数字
	ret = InList(1, []int{1, 2, 3})
	t.Log("ret:", ret)

	// 结构体
	ret = InList(TItem{
		A: 1,
		B: "a",
	}, []TItem{{A: 1, B: "a"}, {A: 2, B: "b"}})
	t.Log("ret:", ret)
}
