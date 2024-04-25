package maps

import "testing"

// func TestSearch(t *testing.T) {
// 	dictionary := map[string]string{"test": "this is just a test"}

// 	got := Search(dictionary, "test")
// 	want := "this is just a test"

// 	assertStrings(t, got, want)
// }

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unkown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		deination := "this is just a test"

		err := dictionary.Add(word, deination)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, deination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		deination := "this is just a test"
		dictionary := Dictionary{word: deination}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, deination)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		deination := "this is just a test"
		dictionary := Dictionary{word: deination}
		newDefination := "new test"

		err := dictionary.Update(word, newDefination)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		deination := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, deination)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
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
