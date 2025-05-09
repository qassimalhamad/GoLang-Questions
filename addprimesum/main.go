package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func toInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			result = result*10 + int(s[i]-'0')
		} else {
			return -1 // invalid input
		}
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n/10 > 0 {
		printNumber(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num := toInt(os.Args[1])
	if num <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printNumber(sum)
	z01.PrintRune('\n')
}
