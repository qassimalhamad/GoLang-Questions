package piscine

import "strconv"

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	return result + strconv.Itoa(count) + string(s[len(s)-1])
}
