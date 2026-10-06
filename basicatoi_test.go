package piscine

import "testing"

func TestBasicAtoi(t *testing.T) {
	for input, want := range map[string]int{
		"12345":         12345,
		"0000000012345": 12345,
		"000000":        0,
		"7":             7,
		"9010":          9010,
	} {
		if got := BasicAtoi(input); got != want {
			t.Errorf("BasicAtoi(%q) = %d, want %d", input, got, want)
		}
	}
}
