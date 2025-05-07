package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] >= 'a' && s[0] <= 'z' {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i-1] == ' ' {
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
		}
	}

	return true
}
