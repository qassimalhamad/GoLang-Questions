package piscine

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	count := 0
	for _, ch := range str {
		if ch != ' ' {
			count++
		}
	}
	if count < 5 {
		return "Invalid Input\n"
	}

	result := ""
	letterCount := 0
	blockCount := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		}

		letterCount++

		if letterCount%6 == 0 {
			continue
		}

		result += string(str[i])
		blockCount++

		if blockCount == 5 {
			result += " "
			blockCount = 0
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
