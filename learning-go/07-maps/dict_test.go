package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		assertValue(t, dict, "test", "this is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is just a test")
		want := "this is just a test"
		assertValue(t, dict, "test", want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dictionary{word: value}
		err := dict.Add("test", "this is just a test")

		assertError(t, err, ErrWordExists)
		assertValue(t, dict, word, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dictionary{word: value}
		new_value := "new value"
		err := dict.Update(word, new_value)

		assertError(t, err, nil)
		assertValue(t, dict, word, new_value)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}
		new_value := "new value"
		err := dict.Update(word, new_value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	value := "this is just a test"
	dict := Dictionary{word: value}
	dict.Delete(word)
	assertNotFound(t, dict, word)
}

func assertNotFound(t testing.TB, dict Dictionary, word string) {
	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertValue(t testing.TB, dict Dictionary, word, value string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, value)
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
