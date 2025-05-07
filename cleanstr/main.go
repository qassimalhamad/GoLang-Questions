package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	result := ""
	inWord := false

	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == ' ' || char == '\t' {
			if inWord {
				result += " "
				inWord = false
			}
		} else {
			result += string(char)
			inWord = true
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	fmt.Print(result + "\n")
}
