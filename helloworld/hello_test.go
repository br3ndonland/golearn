package helloworld

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("say hello with a name argument", func(t *testing.T) {
		got := Hello("Name", "")
		want := "Hello, Name"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when no name is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Nathalie", "French")
		want := "Bonjour, Nathalie"
		assertCorrectMessage(t, got, want)
	})
}
