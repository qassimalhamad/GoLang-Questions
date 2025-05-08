package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	s := str1 + str2
	seen := ""
	count := 0

	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if strings.Contains(seen, ch) {
			continue
		}
		seen += ch

		if strings.Contains(str1, ch) && strings.Contains(str2, ch) {
			continue
		}

		count++
	}
	return count
}
