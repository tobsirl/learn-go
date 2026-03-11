package lostwithoutamap

func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, v := range x {
		result[i] = v * 2
	}
	return result
}
