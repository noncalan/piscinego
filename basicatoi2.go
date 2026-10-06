package piscine

func BasicAtoi2(s string) int {
	n := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0
		}
		n = n*10 + int(digit-'0')
	}
	return n
}
