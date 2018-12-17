package main

import "testing"

var claims = []claim{
	claim{"1", 1, 3, 4, 4},
	claim{"2", 3, 1, 4, 4},
	claim{"3", 5, 5, 2, 2},
	claim{"4", 3, 1, 4, 4},
}

func TestCountOverlaps(t *testing.T) {
	expected := 16
	if result := countOverlaps(claims); result != expected {
		t.Errorf("got: %v want: %v", result, expected)
	}
}

func TestFindUniqueClaimID(t *testing.T) {
	expected := "3"
	if result := findUniqueClaimID(claims); result != expected {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}
