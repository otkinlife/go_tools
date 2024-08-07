package map_tools

import "testing"

func TestSplitMap2List(t *testing.T) {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
		6: "f",
		7: "g",
	}
	chunkSize := 2
	result, err := SplitMap2List(m, chunkSize)
	if err != nil {
		t.Errorf("SplitMap2List() error = %v", err)
		return
	}
	if len(result) != 4 {
		t.Errorf("SplitMap2List() got = %v, want %v", len(result), 4)
	}
	return
}

func TestSplitMap2Map(t *testing.T) {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
		6: "f",
		7: "g",
	}
	chunkSize := 2
	result, err := SplitMap2Map(m, chunkSize)
	if err != nil {
		t.Errorf("SplitMap2Map() error = %v", err)
		return
	}
	if len(result) != 3 {
		t.Errorf("SplitMap2Map() got = %v, want %v", len(result), 3)
	}
	return
}
