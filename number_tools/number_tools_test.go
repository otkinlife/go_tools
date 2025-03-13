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

func TestFormatNumberWithCommas(t *testing.T) {
	r := FormatNumberWithCommas(1234567.89, ",", 2)
	t.Log(r)
	r = FormatNumberWithCommas(1234567.89, ".", 3)
	t.Log(r)
	r = FormatNumberWithCommas(1234567, ",", 0)
	t.Log(r)
}
