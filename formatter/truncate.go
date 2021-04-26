package formatter

import (
	"unicode/utf8"
)

func Truncate(text string, length int) string {
	const ellipsis string = "…"
	if utf8.RuneCountInString(text) > length {
		return text[:length-1] + ellipsis
	}

	return text
}
