package mymod1

import "testing"

func TestHi(t *testing.T) {
	want := "HI12345 !!!"
	if got := Hi(); got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
