package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("outputs Hello, personsName when provided a name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
	t.Run("outputs Hello, World when nothing is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}
	})
}
