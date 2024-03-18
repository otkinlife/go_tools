package color_string

import (
	"testing"
)

func TestColorString(t *testing.T) {
	text := "hello world"
	t.Log(Red(text))
	t.Log(Green(text))
	t.Log(Yellow(text))
}
