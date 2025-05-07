package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	word1 := os.Args[1]
	word2 := os.Args[2]

	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		if word1[i] == word2[j] {
			i++
		}
		j++
	}

	if i == len(word1) {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}
