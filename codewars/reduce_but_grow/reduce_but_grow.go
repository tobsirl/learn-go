package reducebutgrow

func Grow(arr []int) int {
	result := 1
	for _, num := range arr {
		result *= num
	}
	return result
}