package piscine

func RetainFirstHalf(str string) string {
	result := ""
	length := len(str)
	if str == "" {
		return ""
	} else if len(str) == 1 {
		return str
	}

	for i := 0; i < length/2; i++ {
		result += string(str[i])
	}

	return result
}
