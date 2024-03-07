package base64

import "testing"

func TestEncode(t *testing.T) {
	ret := Encode("hello world")
	t.Log("ret:", ret)
}

func TestDecode(t *testing.T) {
	ret, err := Decode("aGVsbG8gd29ybGQ=")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ret:", ret)
}
