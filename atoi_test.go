package piscine

import "testing"

func TestAtoi(t *testing.T) {
	for input, want := range map[string]int{
		"12345":         12345,
		"0000000012345": 12345,
		"012 345":       0,
		"Hello World!":  0,
		"+1234":         1234,
		"-1234":         -1234,
		"++1234":        0,
		"--1234":        0,
		"+-1234":        0,
		"-+1234":        0,
		"":              0,
		"+":             0,
		"-":             0,
		"-00042":        -42,
		"-0":            0,
		"123x":          0,
		" 123":          0,
		"１２３":           0,
	} {
		if got := Atoi(input); got != want {
			t.Errorf("Atoi(%q) = %d, want %d", input, got, want)
		}
	}
}
