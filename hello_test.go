package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello()
	expected := "Hello, world!"
	if actual != expected {
		t.Errorf("actual %q want %q", actual, expected)
	}
}
