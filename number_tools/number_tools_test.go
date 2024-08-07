package number_tools

import "testing"

func TestIntToRoman(t *testing.T) {
	r := IntToRoman(3)
	t.Log(r)
	r = IntToRoman(40)
	t.Log(r)
	r = IntToRoman(9)
	t.Log(r)
}
