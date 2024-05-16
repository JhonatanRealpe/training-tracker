package util

import "testing"

func TestIsEmptyString(t *testing.T) {
	result := IsEmptyString("")
	expected := true

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	result = IsEmptyString("Hello, world!")
	expected = false

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
