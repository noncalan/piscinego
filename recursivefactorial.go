package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb > 20 {
		return 0
	}
	if nb <= 1 {
		return 1
	}

	previous := RecursiveFactorial(nb - 1)
	if previous == 0 {
		return 0
	}
	result := previous * nb
	if result/nb != previous {
		return 0
	}
	return result
}
