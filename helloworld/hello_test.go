package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to someone", func(t *testing.T) {
		got := Hello("Ashmit", "")
		want := "Hello, Ashmit"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Raquel", "Spanish")
		want := "Hola! Raquel"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Annie", "French")
		want := "Bonjour! Annie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
