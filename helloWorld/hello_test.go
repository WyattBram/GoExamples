package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("outputs Hello, personsName when provided a name", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		helper(t, got, want)
	})
	t.Run("outputs Hello, World when nothing is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		helper(t, got, want)
	})
	t.Run("outputs Elodie, World when spanish is passed through", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Elode, World"
		helper(t, got, want)
	})
	t.Run("outputs Elodie, person name when spanish and a persons name is passed through is passed through", func(t *testing.T) {
		got := Hello("Charlie", "Spanish")
		want := "Elode, Charlie"
		helper(t, got, want)
	})
	t.Run("outputs Bonjure, person name when French and a persons name is passed through is passed through", func(t *testing.T) {
		got := Hello("Charlie", "French")
		want := "Bonjure, Charlie"
		helper(t, got, want)
	})
	t.Run("outputs Bonjure, world when French is passed through is passed through", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjure, World"
		helper(t, got, want)
	})
}

func helper(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
