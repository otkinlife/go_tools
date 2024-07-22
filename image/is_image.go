package image

import "strings"

func IsImage(suffix string) bool {
	suffix = strings.ToLower(suffix)
	switch suffix {
	case "jpg", "jpeg", "png", "gif", "bmp", "webp":
		return true
	default:
		return false
	}
}
