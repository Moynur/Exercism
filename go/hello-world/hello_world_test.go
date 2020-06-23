package main

import "testing"

func TestGreet(t *testing.T) {
	if Greet() != "Hello, world!" {
		t.Error("Expected something else >:(")
	}

}
