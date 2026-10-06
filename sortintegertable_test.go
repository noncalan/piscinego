package piscine

import (
	"sort"
	"testing"
)

func TestSortIntegerTable(t *testing.T) {
	for _, input := range [][]int{
		nil,
		{},
		{1},
		{5, 4, 3, 2, 1, 0},
		{-3, 0, 2, 7},
		{3, -1, 3, 0, -5, -1},
		{2, 2, 2},
	} {
		want := append([]int(nil), input...)
		sort.Ints(want)
		SortIntegerTable(input)
		for i := range want {
			if input[i] != want[i] {
				t.Fatalf("got %v, want %v", input, want)
			}
		}
	}
}
