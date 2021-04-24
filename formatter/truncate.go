package formatter

func Truncate(text string, length uint) string {
	if len(text) > int(length) {
		return text[:length-3] + "..."
	}

	return text
}
