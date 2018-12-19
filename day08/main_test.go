package main

import (
	"testing"
)

var input = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}

func TestSumMetadata(t *testing.T) {
	expected := 138
	if r := sumMetadata(input); r != expected {
		t.Errorf("got: %v, want: %v", r, expected)
	}
}

func TestSumNodes(t *testing.T) {
	expected := 66
	if r := sumNodes(input); r != expected {
		t.Errorf("got: %v, want: %v", r, expected)
	}
}
