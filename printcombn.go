package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	digits := make([]rune, n)
	for i := range digits {
		digits[i] = '0' + rune(i)
	}
	for {
		for _, digit := range digits {
			z01.PrintRune(digit)
		}
		i := n - 1
		for i >= 0 && digits[i] == '9'-rune(n-1-i) {
			i--
		}
		if i < 0 {
			break
		}
		digits[i]++
		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}
