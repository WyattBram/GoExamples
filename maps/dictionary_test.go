package maps

import (
	"testing"
)

func TestMap(t *testing.T) {

	dictionary := Dictionary{"Key": "Val"}
	t.Run("Testing when givin a valid key", func(t *testing.T) {

		got, _ := dictionary.Search("Key")
		want := "Val"
		assertTest(t, got, want)

	})

	t.Run("Testing when givin a non-valid key", func(t *testing.T) {

		_, err := dictionary.Search("Testing")

		if err == nil {
			t.Fatal("Expected an error")
		}
		assertError(t, err, NotFound)

	})

}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{}

	dictionary.Add("Key", "Value")

	got, err := dictionary.Search("Key")
	want := "Value"

	if err != nil {
		t.Fatalf("Could not find key, %q", err)
	}

	assertTest(t, got, want)

}

func assertTest(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %q, Want: %q, Given: %q", got, want, "Test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %q, Want: %q", got, want)
	}
}
