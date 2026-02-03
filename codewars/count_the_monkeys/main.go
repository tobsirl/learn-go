package countthemonkeys

func monkeyCount(n int) []int {
	if n <= 0 {
		return []int{}
	}

	monkeys := make([]int, n)
	for i := 0; i < n; i++ {
		monkeys[i] = i + 1
	}
	return monkeys
}