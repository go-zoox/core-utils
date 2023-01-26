package strings

// Slice returns a substring of s.
func Slice(s string, start int, end ...int) string {
	endX := len(s)
	if len(end) > 0 {
		endX = end[0]
	}

	if start < 0 {
		start = 0
	}

	if endX < 0 {
		endX = len(s) + endX
	}

	if endX > len(s) {
		endX = len(s)
	}

	if start > endX {
		return ""
	}

	return s[start:endX]
}
