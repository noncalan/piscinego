package piscine

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintComb(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	stdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = stdout }()
	PrintComb()
	w.Close()
	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(string(output), "\n") {
		t.Fatal("missing final newline")
	}
	combinations := strings.Split(strings.TrimSuffix(string(output), "\n"), ", ")
	if len(combinations) != 120 {
		t.Fatalf("got %d combinations, want 120", len(combinations))
	}
	previous := ""
	for _, combination := range combinations {
		if len(combination) != 3 || combination[0] < '0' || combination[2] > '9' ||
			combination[0] >= combination[1] || combination[1] >= combination[2] || combination <= previous {
			t.Fatalf("invalid or unordered combination: %q", combination)
		}
		previous = combination
	}
}
