/*
Dependency injection

`fmt.Printf` prints text by calling `fmt.Fprintf` with `os.Stdout` as an argument.
`os.Stdout` implements the `io.Writer` interface.
`bytes.Buffer` also implements the `io.Writer` interface.
This test "injects" a `bytes.Buffer` in place of `os.Stdout`
for more control over the contents, like a form of mocking.
*/
package di

import (
	"bytes"
	"testing"
)

func assertCorrectValue(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "World")
	got := buffer.String()
	want := "Hello, World!"
	assertCorrectValue(t, got, want)
}
