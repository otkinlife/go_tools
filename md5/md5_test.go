package md5

import "testing"

func TestFileGenerate(t *testing.T) {
	m, err := FileGenerate("readme.md")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("md5:", m)
}

func TestStringGenerate(t *testing.T) {
	m := StringGenerate("hello world")
	t.Log("md5:", m)
}
