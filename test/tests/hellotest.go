package test

import (
	"alderaan/test"
	"testing"
)

func TestHello(t *testing.T) {
	got := test.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
