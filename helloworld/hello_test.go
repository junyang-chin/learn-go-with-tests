package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jun", "Default")
		want := "Hello, Jun"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print Hello World when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jun", "Spanish")
		want := "Hola, Jun"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jun", "French")
		want := "Bonjour, Jun"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
