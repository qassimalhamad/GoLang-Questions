package piscine

func HashCode(dec string) string {
	result := ""
	for i := 0; i < len(dec); i++ {
		temp := (int(dec[i]) + len(dec)) % 127

		if temp < 33 {
			temp = temp + 33
		}

		result += string(temp)
	}
	return result
}
