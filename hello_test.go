package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		recipient := "Ömer"
		want, got := fmt.Sprintf("Hello, %s", recipient), Hello(recipient)
		if want != got {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		want, got := "Hello, World", Hello("")

		if want != got {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
}
