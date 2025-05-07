package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 4 {
		return
	}
	word := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]
	result := ""

	for i := 0; i < len(word); i++ {
		if string(word[i]) == search {
			result += replace
		} else {
			result += string(word[i])
		}
	}
	fmt.Print(result + "\n")
}
