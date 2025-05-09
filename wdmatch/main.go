package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	i := 0

	for j := 0; j < len(s2); j++ {
		if i < len(s1) && s1[i] == s2[j] {
			i++
		}
	}

	if i == len(s1) {
		fmt.Print(s1)
	}
	fmt.Println()
}
