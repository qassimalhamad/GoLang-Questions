package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	var result [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}

		result = append(result, slice[i:end])
	}
	fmt.Println(result)
}
