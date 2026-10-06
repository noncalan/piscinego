package piscine

func Atoi(s string) int {
	sign := 1
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	n := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0
		}
		n = n*10 + int(digit-'0')
	}
	return sign * n
}
