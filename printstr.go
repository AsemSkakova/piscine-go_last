package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	char := []rune(str)
	for index := range char {
		z01.PrintRune(char[index])
	}
	z01.PrintRune(10)
}
