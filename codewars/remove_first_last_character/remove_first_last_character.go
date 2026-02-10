package removefirstlastcharacter

func RemoveChar(word string) string {
	if len(word) <= 2 {
		return ""
	}
	return word[1 : len(word)-1]
}