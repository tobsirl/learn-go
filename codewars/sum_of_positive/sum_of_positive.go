package sumofpositive

func PositiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}

	return sum
}