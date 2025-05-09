package piscine

import (
	"math"
	"strconv"
)

func NotDecimal(dec string) string {
	// If input is empty, return newline
	if dec == "" {
		return "\n"
	}

	// Look for the decimal point
	dotIndex := -1
	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			dotIndex = i
			break
		}
	}

	// If there's no decimal point, return as is
	if dotIndex == -1 {
		return dec + "\n"
	}

	// Check if everything after the dot is just zeros
	onlyZeros := true
	for i := dotIndex + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			onlyZeros = false
			break
		}
	}
	if onlyZeros {
		return dec + "\n"
	}

	// Try to convert the string to a float
	number, err := strconv.ParseFloat(dec, 64)
	if err != nil {
		// If it's not a valid number, return as is
		return dec + "\n"
	}

	// Calculate how many digits are after the decimal point
	digitsAfterDot := len(dec) - dotIndex - 1

	// Shift decimal to the right by multiplying by 10^n
	shifted := number * math.Pow(10, float64(digitsAfterDot))

	// Convert final result to int and return as string
	return strconv.Itoa(int(shifted)) + "\n"
}
