package piscine

import (
	"os"
	"strings"
	"testing"
)

func TestPrintComb2(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "output")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	stdout := os.Stdout
	os.Stdout = f
	defer func() { os.Stdout = stdout }()
	PrintComb2()
	output, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(string(output), "\n") {
		t.Fatal("missing final newline")
	}
	pairs := strings.Split(strings.TrimSuffix(string(output), "\n"), ", ")
	if len(pairs) != 4950 {
		t.Fatalf("got %d pairs, want 4950", len(pairs))
	}
	previous := ""
	for _, pair := range pairs {
		if len(pair) != 5 || pair[2] != ' ' || pair[:2] >= pair[3:] || pair <= previous {
			t.Fatalf("invalid or unordered pair: %q", pair)
		}
		for _, digit := range pair[:2] + pair[3:] {
			if digit < '0' || digit > '9' {
				t.Fatalf("invalid digit in %q", pair)
			}
		}
		previous = pair
	}
}
