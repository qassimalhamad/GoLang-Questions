package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1] + os.Args[2]
	seen := ""

	for i := 0; i < len(input); i++ {
		char := input[i]
		repeated := false

		for j := 0; j < len(seen); j++ {
			if seen[j] == char {
				repeated = true
				break
			}
		}
		if !repeated {
			seen += string(char)
			z01.PrintRune(rune(char))
		}
	}
	z01.PrintRune('\n')
}
