package dictionary_test

import (
	"testing"

	"github.com/MarkMoelter/learn-golang/dictionary"
)

func TestSearch(t *testing.T) {
	dictionary_ := dictionary.Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary_.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary_.Search("unknown")

		if err != dictionary.ErrNotFound {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, dictionary.ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := dictionary.Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary_ := dictionary.Dictionary{word: definition}

		err := dictionary_.Add(word, "new test")

		assertError(t, err, dictionary.ErrWordExists)
		assertDefinition(t, dictionary_, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := dictionary.Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary_ := dictionary.Dictionary{}

		err := dictionary_.Update(word, definition)

		assertError(t, err, dictionary.ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary_ := dictionary.Dictionary{word: definition}

		err := dictionary_.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary_.Search(word)
		assertError(t, err, dictionary.ErrNotFound)
	})
	t.Run("delete non-existent word", func(t *testing.T) {
		word := "test"
		dictionary_ := dictionary.Dictionary{}

		err := dictionary_.Delete(word)

		assertError(t, err, dictionary.ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary dictionary.Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
