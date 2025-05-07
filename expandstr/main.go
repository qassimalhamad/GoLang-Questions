package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	result := ""
	inWord := false
	spaceNeeded := false

	for i := 0; i < len(input); i++ {
		char := input[i]

		if char == ' ' || char == '\t' {
			if inWord {
				spaceNeeded = true
				inWord = false
			}
		} else {
			if spaceNeeded {
				result += "   "
				spaceNeeded = false
			}
			result += string(char)
			inWord = true
		}
	}

	if result != "" {
		fmt.Println(result)
	}
}
