package main

import (
	"testing"
)

func TestAddSix(t *testing.T) {
	if AddSix(2) != 8 {
		t.Error("Expected 2 + 6 to equal 8")
	}
}
