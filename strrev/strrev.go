package piscine

func StrRev(s string) string {
	var k int
	var a int
	reverse := []byte(s)
	for index, word := range s {
		if word == 233 {
			a = -1
		}
		k = index + a
	}

	j := 0
	for i := k; i >= 0; i-- {
		reverse[i] = byte(s[j])
		j++
	}
	out := string(reverse)
	return out
}
