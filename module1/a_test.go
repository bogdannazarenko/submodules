package module1

import "testing"

func TestHi(t *testing.T) {
	want := "HI!!! 123"
	if got := Hi(); got != want {
		t.Fatalf("got: %s, want: %s", got, want)
	}
}
