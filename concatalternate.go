package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	var longer []int
	var shorter []int
	var result []int

	if len(slice1) > len(slice2) {
		longer = slice1
		shorter = slice2
	} else if len(slice2) > len(slice1) {
		longer = slice2
		shorter = slice1
	} else {
		for i := 0; i < len(slice1); i++ {
			result = append(result, slice1[i], slice2[i])
		}
	}

	for i := 0; i < len(shorter); i++ {
		result = append(result, longer[i], shorter[i])
	}

	for i := len(shorter); i < len(longer); i++ {
		result = append(result, longer[i])
	}
	return result
}
