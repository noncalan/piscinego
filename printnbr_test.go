package piscine

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPrintNbr(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "output")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	stdout := os.Stdout
	os.Stdout = f
	defer func() { os.Stdout = stdout }()
	max := int(^uint(0) >> 1)
	values := []int{-123, 0, 123, -max - 1, max, -10, -9, -1, 1, 9, 10, 100, -100}
	var want strings.Builder
	for _, n := range values {
		PrintNbr(n)
		want.WriteString(strconv.Itoa(n))
	}
	got, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want.String() {
		t.Fatalf("got %q, want %q", got, want.String())
	}
}
