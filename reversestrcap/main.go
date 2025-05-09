package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for a := 1; a < len(os.Args); a++ {
		arg := os.Args[a]

		for i := 0; i < len(arg); i++ {
			char := arg[i]
			isLetter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
			isLast := i == len(arg)-1 || arg[i+1] == ' ' || arg[i+1] == '\t'

			if isLast && isLetter {
				z01.PrintRune(rune(toUpper(char)))
			} else {
				z01.PrintRune(rune(toLower(char)))
			}
		}
		if len(arg)-1 > a {
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}
