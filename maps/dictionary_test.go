package maps

import (
	"testing"
)

func assertCorrectValue(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("error expected")
	}
	gotErr := got.Error()
	wantErr := want.Error()
	if gotErr != wantErr {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMap(t *testing.T) {
	mapKey := "test"
	mapValue := "this is just a test"
	mapInstance := map[string]string{mapKey: mapValue}
	t.Run("key in map", func(t *testing.T) {
		got, inMap := mapInstance[mapKey]
		want := mapValue
		if got == "" || !inMap {
			t.Fatal("value expected")
		}
		assertCorrectValue(t, got, want)
	})
	t.Run("key not in map", func(t *testing.T) {
		unknownKey := "unknown"
		got, inMap := mapInstance[unknownKey]
		if got != "" || inMap {
			t.Fatal("no value expected")
		}
	})

}

func TestGet(t *testing.T) {
	dictionaryKey := "test"
	dictionaryValue := "this is just a test"
	dictionary := Dictionary{dictionaryKey: dictionaryValue}
	t.Run("key in map", func(t *testing.T) {
		got, _ := dictionary.Get(dictionaryKey)
		want := dictionaryValue
		assertCorrectValue(t, got, want)
	})
	t.Run("key not in map", func(t *testing.T) {
		unknownKey := "unknown"
		_, err := dictionary.Get(unknownKey)
		if err == nil {
			t.Fatal("error expected")
		}
		want := dictionary.KeyError(unknownKey)
		assertError(t, err, want)
	})
}

func TestSet(t *testing.T) {
	dictionaryKey := "test"
	dictionaryValue := "this is just a test"
	dictionary := Dictionary{}
	dictionary.Set(dictionaryKey, dictionaryValue)
	got, err := dictionary.Get(dictionaryKey)
	if err != nil {
		t.Fatal("should find new dictionary key", err)
	}
	assertCorrectValue(t, got, dictionaryValue)
}

func TestAdd(t *testing.T) {
	dictionaryKey := "test"
	dictionaryValue := "this is just a test"
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(dictionaryKey, dictionaryValue)
		if err != nil {
			t.Fatal("should not raise an error")
		}
		assertCorrectValue(t, dictionary[dictionaryKey], dictionaryValue)
	})
	t.Run("existing key", func(t *testing.T) {
		newDictionaryValue := "new value"
		dictionary := Dictionary{dictionaryKey: dictionaryValue}
		err := dictionary.Add(dictionaryKey, newDictionaryValue)
		want := dictionary.ExistsError(dictionaryKey)
		if err == nil {
			t.Fatal("should raise an error")
		}
		if err != nil {
			assertError(t, err, want)
		}
		assertCorrectValue(t, dictionary[dictionaryKey], dictionaryValue)
	})
}
