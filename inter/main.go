package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	seen := ""

	for i := 0; i < len(s1); i++ {
		char := s1[i]

		repeated := false
		for j := 0; j < len(seen); j++ {
			if char == seen[i] {
				repeated = true
				break
			}
		}

		if repeated {
			continue
		}
		found := false
		for j := 0; j < len(s2); j++ {
			if char == s2[j] {
				found = true
				break
			}
		}
		if found {
			z01.PrintRune(rune(char))
			seen += string(char)
		}
	}
	z01.PrintRune('\n')
}
