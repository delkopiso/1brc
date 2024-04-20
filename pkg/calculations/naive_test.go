package calculations_test

import (
	"bytes"
	"os"
	"testing"

	"1brc/pkg/calculations"
)

func TestCalculateNaive(t *testing.T) {
	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = file.Close()
	})

	_, _ = file.WriteString("foo;1\nbar;1\nbaz;1\n")
	_, _ = file.WriteString("bar;2\nbaz;2\nfoo;2\n")
	_, _ = file.WriteString("baz;3\nfoo;3\nbar;3\n")

	buf := new(bytes.Buffer)
	calculations.CalculateNaive(file.Name(), buf)

	expected := "{bar=1.0/2.0/3.0, baz=1.0/2.0/3.0, foo=1.0/2.0/3.0}"
	if buf.String() != expected {
		t.Errorf("expected %s, but got %s", expected, buf.String())
	}
}
