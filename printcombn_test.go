package piscine

import (
	"os"
	"strings"
	"testing"
)

func TestPrintCombN(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "output")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	stdout := os.Stdout
	os.Stdout = f
	defer func() { os.Stdout = stdout }()
	for n := 1; n < 10; n++ {
		PrintCombN(n)
	}
	output, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) != 10 || lines[9] != "" {
		t.Fatal("expected nine newline-terminated lines")
	}
	count := 1
	for n := 1; n < 10; n++ {
		count = count * (11 - n) / n
		combinations := strings.Split(lines[n-1], ", ")
		if len(combinations) != count {
			t.Fatalf("n=%d: got %d combinations, want %d", n, len(combinations), count)
		}
		previous := ""
		for _, combination := range combinations {
			if len(combination) != n || combination <= previous {
				t.Fatalf("n=%d: invalid or unordered combination %q", n, combination)
			}
			last := rune('0' - 1)
			for _, digit := range combination {
				if digit <= last || digit > '9' {
					t.Fatalf("n=%d: invalid digits in %q", n, combination)
				}
				last = digit
			}
			previous = combination
		}
	}
}
