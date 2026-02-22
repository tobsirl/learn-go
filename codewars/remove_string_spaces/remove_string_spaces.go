package removestringspaces

func NoSpace(word string) string {
	result := ""
	for _, char := range word {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}