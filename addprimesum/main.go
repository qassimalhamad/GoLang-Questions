package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	input := os.Args[1]
	num, err := strconv.Atoi(input)

	if err != nil {
		return
	}

	sum := 0

	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Print(sum)
	fmt.Print("\n")
}
