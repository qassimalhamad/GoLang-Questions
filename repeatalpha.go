package piscine

func RepeatAlpha(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'a' && ch <= 'z' {
			for i := 'a'; i <= rune(ch); i++ {
				result += string(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			for i := 'A'; i <= rune(ch); i++ {
				result += string(ch)
			}
		} else {
			result += string(ch)
		}
	}
	return result
}
