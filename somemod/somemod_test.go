package somemod

import "testing"

func TestSomeMod(t *testing.T) {
	if 2 != 2 {
		t.Fatalf("If this doesn't pass the universe has fundamentally changed. Try a reboot.")
	}
}
