package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mark")
	want := "Hello, Mark"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
