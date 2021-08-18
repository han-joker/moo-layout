package dftm

import "testing"

func TestString(t *testing.T) {
	if String("", "default") != "default" {
		t.Error("string default error")
	}
}
