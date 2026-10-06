package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n <= -10 {
			PrintNbr(-(n / 10))
		}
		n = -(n % 10)
	} else if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune('0' + rune(n%10))
}
