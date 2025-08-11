package main

import (
	"testing"
)

func TestAddSix(t *testing.T) {
	if AddSix(2) != 8 {
		t.Error("Expected 2 + 6 to equal 8")
	}
	if AddSix(3.3) != 9.3 {
		t.Error("Expected 3.3 + 6 to equal 9.3")
	}
}
