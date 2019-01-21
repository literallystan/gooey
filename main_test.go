package main

import "testing"

func TestMain(t *testing.T) {
	if 2 != 2 {
		t.Fatalf("If this fails the universe has fundamentally changed. Try a reboot.")
	}
}
