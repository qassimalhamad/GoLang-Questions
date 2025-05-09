package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	seen := ""

	for i := 0; i < len(str1); i++ {
		char := str1[i]

		repeated := false
		for j := 0; j < len(seen); j++ {
			if char == seen[j] {
				repeated = true
				break
			}
		}
		if repeated {
			continue
		}
		found := false
		for k := 0; k < len(str2); k++ {
			if char == str2[k] {
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
