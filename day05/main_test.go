package main

import "testing"

func TestGetPolimerLength(t *testing.T) {
	data := "dabAcCaCBAcCcaDA"
	expected := 10
	if result := getPolimerLength(data); result != expected {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestGetImprovedPolimerLength(t *testing.T) {
	data := "dabAcCaCBAcCcaDA"
	expected := 4
	if result := getImprovedPolimerLength(data); result != expected {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
