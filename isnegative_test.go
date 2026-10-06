package piscine

import (
	"io"
	"os"
	"testing"
)

func TestIsNegative(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	stdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = stdout }()
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
	w.Close()
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "F\nF\nT\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}
