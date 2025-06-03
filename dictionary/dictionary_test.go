package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should not receive error but got one")
		}
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown world", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}

		err := dictionary.Add(term, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, term, definition)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertDefinition(t *testing.T, dictionary Dictionary, term, definition string) {
	t.Helper()
	got, err := dictionary.Search(term)
	if err != nil {
		t.Fatal("should not receive error but got one")
	}
	assertStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition, definition: term}
		newDefinition := "new definition"
		err := dictionary.Update(term, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, term, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(term, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrWordNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
