package piscine

import "testing"

func TestStrRev(t *testing.T) {
	for input, want := range map[string]string{
		"":             "",
		"a":            "a",
		"abc":          "cba",
		"Hello World!": "!dlroW olleH",
		"Привет🌍":      "🌍тевирП",
	} {
		if got := StrRev(input); got != want {
			t.Errorf("StrRev(%q) = %q, want %q", input, got, want)
		}
	}
}
