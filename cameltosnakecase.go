package piscine

func CamelToSnakeCase(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		if s[i] > 'A' && s[i] < 'Z' && s[i+1] > 'A' && s[i+1] < 'Z' {
			return s
		}

		if s[i] > 'A' && s[i] < 'Z' && i != 0 {
			result += "_" + string(s[i])
		} else {
			result += string(s[i])
		}
	}
	return result
}
