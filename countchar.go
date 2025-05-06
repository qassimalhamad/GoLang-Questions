package piscine

func CountChar(str string, c rune) int {
	counter := 0
	for i := 0; i < len(str); i++ {
		if str[i] == byte(c) {
			counter++
		}
	}
	return counter
}
