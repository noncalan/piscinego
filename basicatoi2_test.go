package piscine

import "testing"

func TestBasicAtoi2(t *testing.T) {
	for input, want := range map[string]int{
		"12345":         12345,
		"0000000012345": 12345,
		"012 345":       0,
		"Hello World!":  0,
		"":              0,
		"000000":        0,
		"9":             9,
		"123x":          0,
		"+123":          0,
		"-123":          0,
		"12/3":          0,
		"12:3":          0,
		"１２３":           0,
	} {
		if got := BasicAtoi2(input); got != want {
			t.Errorf("BasicAtoi2(%q) = %d, want %d", input, got, want)
		}
	}
}
