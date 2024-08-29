package wc

func CountBytes(content []byte) int {
	return len(content)
}

func CountLines(content []byte) int {
	var count int
	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			count++
		}
	}
	return count
}

func isWhitespace(b byte) bool {
	return b == ' ' ||
		b == '\n' ||
		b == '\r' ||
		b == '\t'
}

func CountWords(content []byte) int {
	var count int
	var inWord bool
	for i := 0; i < len(content); i++ {
		if isWhitespace(content[i]) && inWord {
			inWord = false
			count++
		} else if !isWhitespace(content[i]) {
			inWord = true
		}
	}
	return count
}
