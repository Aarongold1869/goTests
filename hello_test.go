package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aaron")
	want := "Hello, Aaron"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
