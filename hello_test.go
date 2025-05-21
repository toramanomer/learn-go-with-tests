package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	recipient := "Ömer"
	want, got := fmt.Sprintf("Hello, %s", recipient), Hello(recipient)
	if want != got {
		t.Errorf("wanted %q, got %q", want, got)
	}
}
