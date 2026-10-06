package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	maxInt := int(^uint(0) >> 1)
	for !IsPrime(nb) {
		if nb > maxInt-2 {
			return 0
		}
		nb += 2
	}
	return nb
}
