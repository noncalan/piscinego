package piscine

import (
	"io"
	"os"
	"testing"
)

func TestPrintStr(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()
	stdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = stdout }()
	PrintStr("")
	PrintStr("Hello World! Привет 🌍")
	w.Close()
	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "Hello World! Привет 🌍" {
		t.Fatalf("unexpected output: %q", got)
	}
}
