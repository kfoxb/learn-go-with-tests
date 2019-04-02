package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Kyle")
	want := "Hello, Kyle"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
