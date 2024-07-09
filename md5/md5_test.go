package md5

import "testing"

func TestFileGenerate(t *testing.T) {
	m, err := FileGenerate("/Users/admin/Downloads/feae19dc-64aa-4e9a-9bdf-8585f3d4a13c.mp4")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("md5:", m)
}

func TestStringGenerate(t *testing.T) {
	m := StringGenerate("hello world")
	t.Log("md5:", m)
}
