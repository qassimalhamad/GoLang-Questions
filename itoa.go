package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false

	if n < 0 {
		isNegative = true
	}

	if isNegative {
		n = n * -1
	}

	result := ""
	for n > 0 {
		result = string(n%10+'0') + result
		n = n / 10
	}

	if isNegative {
		result = "-" + result
	}
	return result
}
