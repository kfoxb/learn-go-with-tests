package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("say hello to people by name", func(t *testing.T) {
		got := Hello("Kyle", "")
		want := "Hello, Kyle"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when given an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("works in Spanish", func(t *testing.T) {
		got := Hello("Kyle", "Spanish")
		want := "Hola, Kyle"
		assertCorrectMessage(t, got, want)
	})

	t.Run("works in French", func(t *testing.T) {
		got := Hello("Kyle", "French")
		want := "Bonjour, Kyle"
		assertCorrectMessage(t, got, want)
	})
}
