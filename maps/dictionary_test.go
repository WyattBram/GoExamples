package maps

import (
	"testing"
)

func TestMap(t *testing.T) {

	dictionary := Dictionary{"Key": "Val"}

	got := dictionary.Search("Key")
	want := "Val"
	assertTest(t, got, want)
}

func assertTest(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %q, Want: %q, Given: %q", got, want, "Test")
	}

}
