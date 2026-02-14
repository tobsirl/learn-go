package smallestintegerfinder

func SmallestIntegerFinder(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Return 0 for an empty slice
	}

	smallest := numbers[0]
	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}