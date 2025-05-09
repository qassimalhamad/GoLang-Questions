package piscine

func SaveAndMiss(arg string, num int) string {
	if num < 0 || num == 0 {
		return arg
	}

	save := true
	count := 0
	result := ""
	for i := 0; i < len(arg); i++ {
		if save {
			result += string(arg[i])
		}
		count++

		if count == num {
			save = !save
			count = 0
		}
	}
	return result
}
