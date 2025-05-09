package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	result := []int{}

	for i := 0; i < len(slice1); i++ {
		result = append(result, slice1[i])
	}

	for i := 0; i < len(slice2); i++ {
		result = append(result, slice2[i])
	}

	return result
}
