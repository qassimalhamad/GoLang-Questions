package piscine

func FirstWord(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			break
		} else {
			result += string(s[i])
		}
	}
	return result + "\n"
}
