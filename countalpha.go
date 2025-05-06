package piscine

func CountAlpha(s string) int {
	counter := 0
	for v := 0; v < len(s); v++ {
		if s[v] > 'A' && s[v] < 'Z' || s[v] > 'a' && s[v] < 'z' {
			counter++
		}
	}
	return counter
}
