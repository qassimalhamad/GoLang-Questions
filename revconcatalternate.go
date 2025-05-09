package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}

	i := len(slice1) - 1
	j := len(slice2) - 1

	for len(slice1) > len(slice2) && i > j {
		result = append(result, slice1[i])
		i--
	}
	for len(slice2) > len(slice1) && j > i {
		result = append(result, slice2[j])
		j--
	}

	for i >= 0 && j >= 0 {
		result = append(result, slice1[i])
		result = append(result, slice2[j])
		i--
		j--
	}

	return result
}
