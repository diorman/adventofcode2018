package main

import "testing"

// func TestCartTrackFindFirstCrash(t *testing.T) {
// 	track, err := parseInputFile("input_test.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	expected := coordinate{7, 3}
// 	if position := track.findFirstCrash(); position != expected {
// 		t.Errorf("got: %v, expected: %v", position, expected)
// 	}
// }

func TestFindLastCartLocation(t *testing.T) {
	track, err := parseInputFile("input_test_2.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected := coordinate{6, 4}
	if position := track.findLastCartLocation(); position != expected {
		t.Errorf("got: %v, expected: %v", position, expected)
	}
}
