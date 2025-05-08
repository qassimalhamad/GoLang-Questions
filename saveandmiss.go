package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""
	save := true
	counter := 0
	for i := 0; i < len(arg); i++ {

		if save {
			result += string(arg[i])
		}
		counter++

		if num == counter {
			save = !save
			counter = 0
		}

	}
	return result
}
