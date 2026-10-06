package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var a rune
	for a = 0; a < 99; a++ {
		for b := a + 1; b < 100; b++ {
			if a != 0 || b != 1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(a/10 + '0')
			z01.PrintRune(a%10 + '0')
			z01.PrintRune(' ')
			z01.PrintRune(b/10 + '0')
			z01.PrintRune(b%10 + '0')
		}
	}
	z01.PrintRune('\n')
}
