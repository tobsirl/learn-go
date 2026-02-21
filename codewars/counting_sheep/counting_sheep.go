package countingsheep

func CountSheep(numbers []bool) int {
	if numbers == nil {
		return 0
	}
	count := 0
	for _, present := range numbers {
		if present {
			count++
		}
	}
	return count
}