package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		if err != nil {
			t.Fatal("should have added the word:", err)
		}

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		word := "test"
		definition := "this is just"
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "word"
		definition := "This is just a test"
		updated := "This is updated test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, updated)

		if err != nil {
			t.Fatal("should have updated the word", err)
		}

		assertDefinition(t, dictionary, word, updated)
	})

	t.Run("new word", func(t *testing.T) {
		word := "word"
		definition := "This is just a test"
		updated := "This is updated test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update("unknown", updated)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is the test"}
		err := dictionary.Delete(word)

		if err != nil {
			t.Fatal("should have deleted the word:", err)
		}

		_, err = dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("got %v, want %v", err, ErrNotFound)
		}
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
