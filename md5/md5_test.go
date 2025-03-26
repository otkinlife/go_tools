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

func TestFileUrlGenerate(t *testing.T) {
	m, err := FileUrlGenerate("https://bi-material-center.oss-cn-beijing.aliyuncs.com/materials_center/videos/3c491c3372ac383965e9a8202eb4e3a6/3c491c3372ac383965e9a8202eb4e3a6.mp4", "3c491c3372ac383965e9a8202eb4e3a6.mp4")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("md5:", m)
}
