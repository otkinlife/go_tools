package img

import "strings"

func IsImage(suffix string) bool {
	suffix = strings.ToLower(suffix)
	switch suffix {
	case Jpg, Jpeg, Png, Gif, Bmp, Webp:
		return true
	default:
		return false
	}
}
