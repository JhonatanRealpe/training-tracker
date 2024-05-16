package util

import "testing"

func TestStrinToInt(t *testing.T) {
	result, _ := StrinToInt("10")
	expected := 10

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestIntToStrin(t *testing.T) {
	result := IntToStrin(10)
	expected := "10"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestStrinToInt_Error(t *testing.T) {
	_, err := StrinToInt("abc")

	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
