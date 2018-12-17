package main

import "testing"

func TestSort(t *testing.T) {
	steps, err := parseInputFile("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected := "CABDFE"
	if r := steps.sort().String(); r != expected {
		t.Errorf("got: %v, want: %v", r, expected)
	}
}

func TestDuration(t *testing.T) {
	steps, err := parseInputFile("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected := 15
	if r := steps.duration(2, 0); r != expected {
		t.Errorf("got: %v, want: %v", r, expected)
	}
}
