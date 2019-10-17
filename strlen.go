package piscine

func StrLen(str string) int {
	char := []rune(str)
	num := 0
	for index := range char {
		num = index + 1
	}
	return num
}
