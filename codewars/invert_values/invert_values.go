package invertvalues

func Invert(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = -v
	}
	return result
}
