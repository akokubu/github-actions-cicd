package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello"
	got := Greeting()
	if want != got {
		t.Errorf("got %q wat %q", got, want)
	}
}
