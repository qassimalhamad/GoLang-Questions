package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	i := 2
	for num > 1 {
		if num%i == 0 {
			fmt.Print(i)
			num /= i
			if num > 1 {
				fmt.Print("*")
			}
		} else {
			i++
		}
	}
	fmt.Println()
}
