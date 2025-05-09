package piscine

import "strings"

func FifthAndSkip(str string) string {
	// If string is empty
	if len(str) == 0 {
		return "\n"
	}

	// Remove all spaces
	str = strings.ReplaceAll(str, " ", "")

	// Check if enough characters
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(str); i++ {
		// Skip every 6th character
		if (i+1)%6 == 0 {
			continue
		}

		result += string(str[i])
		count++

		// Add space every 5 characters
		if count%5 == 0 {
			result += " "
		}
	}

	// Remove trailing space if any
	if strings.HasSuffix(result, " ") {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
