package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Arvin")
	want := "Hello, Arvin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}