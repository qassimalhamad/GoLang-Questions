package piscine

func ThirdTimeIsACharm(str string) string {
	result := ""
	for i := 2; i < len(str); i = i + 3 {
		result += string(str[i])
	}
	return result + "\n"
}
