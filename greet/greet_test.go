package greet_test

import (
	"bytes"
	"testing"

	"github.com/MarkMoelter/learn-golang/greet"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	greet.Greet(&buf, "Chris")

	got := buf.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
