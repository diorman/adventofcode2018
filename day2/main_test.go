package main

import "testing"

func TestCalculateChecksum(t *testing.T) {
	ids := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	expected := 12
	if result := calculateChecksum(ids); result != expected {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestFindCommonLetters(t *testing.T) {
	ids := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	expected := "fgij"
	if result := getCommonLetters(ids); result != expected {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
