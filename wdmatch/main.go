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

	i := 0
	for j := 0; j < len(str2); j++ {
		if i < len(str1) && str1[i] == str2[j] {
			i++
		}
	}

	if i == len(str1) {
		for k := 0; k < len(str1); k++ {
			z01.PrintRune(rune(str1[k]))
		}
	}
	z01.PrintRune('\n')
}
