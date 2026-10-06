package piscine

func SortIntegerTable(table []int) {
	// ponytail: O(n²) insertion sort; use heapsort if large inputs matter.
	for i := 1; i < len(table); i++ {
		for j := i; j > 0 && table[j] < table[j-1]; j-- {
			table[j], table[j-1] = table[j-1], table[j]
		}
	}
}
