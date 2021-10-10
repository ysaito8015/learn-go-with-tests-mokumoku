package iteration

import "testing"

func TestRepeat(t *testing.T) {
	output := Repeat("a")
	expected := "aaaaa"

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}
