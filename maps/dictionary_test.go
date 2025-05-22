package maps

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		def := "this is a test definition"
		err := dictionary.Add("test", def)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test definition"
		dictionary := Dictionary{word: def}
		err := dictionary.Add("test", def)

		assertError(t, err, ErrWordExists)
	})
}

func TestDelete(t *testing.T) {
	var (
		word = "test"
		def  = "test def"
		dict = Dictionary{word: def}
	)
	dict.Delete(word)
	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("update existing", func(t *testing.T) {
		var (
			word       = "test"
			def        = "test def"
			dictionary = Dictionary{word: def}
			newDef     = "new def"
		)
		err := dictionary.Update(word, newDef)
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDef)
	})

	t.Run("update non existing", func(t *testing.T) {
		var (
			dictionary = Dictionary{}
			word       = "test"
			newDef     = "new def"
		)
		err := dictionary.Update(word, newDef)
		assertError(t, err, ErrNotFound)
	})

}

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		var (
			want     = "this is just a test"
			got, err = dictionary.Search("test")
		)
		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		var (
			want     = ""
			got, err = dictionary.Search("test")
		)

		assertStrings(t, got, want)
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted: %q, got: %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("expected error: %#v, got: %#v", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}
