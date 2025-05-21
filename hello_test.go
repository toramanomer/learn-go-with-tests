package main

import (
	"fmt"
	"strings"
	"testing"
)

func assertCorretMessage(t testing.TB, want, got string) {
	// This is needed to tell the test suite that this method is helper.
	// So when it fails, the line number reported will be in our function call
	// rather than this function.
	t.Helper()
	if want != got {
		t.Errorf("wanted %q, got %q", want, got)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		recipient := "Ömer"
		want, got := fmt.Sprintf("Hello, %s", recipient), Hello(recipient, "")
		assertCorretMessage(t, want, got)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		want, got := "Hello, World", Hello("", "")
		assertCorretMessage(t, want, got)
	})

	t.Run("in Spanish", func(t *testing.T) {
		want, got := "Hola, Ömer", Hello("Ömer", "Spanish")
		assertCorretMessage(t, want, got)
	})

	t.Run("in French", func(t *testing.T) {
		want, got := "Bonjour, Ömer", Hello("Ömer", "French")
		assertCorretMessage(t, want, got)
	})
}

func TestGetLanguagePrefix(t *testing.T) {
	t.Run("returns correct prefix regardless of case", func(t *testing.T) {
		want, got := englishHelloPrefix, GetLanguagePrefix(strings.ToUpper(string(English)))
		assertCorretMessage(t, want, got)
	})

	t.Run("returns english if language is not found", func(t *testing.T) {
		want, got := GetLanguagePrefix(string(English)), GetLanguagePrefix("something_that_does_not_exist")
		assertCorretMessage(t, want, got)
	})
}
