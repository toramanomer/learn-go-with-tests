package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	Greet(buffer, "Omer")

	var (
		got  = buffer.String()
		want = "Hello, Omer"
	)
	if want != got {
		t.Errorf("wanted: %q, got: %q", want, got)
	}
}
