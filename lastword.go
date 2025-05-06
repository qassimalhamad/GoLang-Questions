package piscine

func LastWord(s string) string {
	result := ""

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' {
			if result != "" {
				break
			}
		} else {
			result = string(s[i]) + result
		}
	}
	return result + "\n"
}
