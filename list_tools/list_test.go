package list_tools

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

func TestFindDifferenceNotInSlice2(t *testing.T) {
	// 字符串
	ret := FindDifferenceNotInSlice2([]string{"a", "b", "c"}, []string{"a", "c"})
	t.Log("ret:", ret)
	// 数字
	ret2 := FindDifferenceNotInSlice2([]int{1, 2, 3}, []int{1, 3})
	t.Log("ret:", ret2)
}
