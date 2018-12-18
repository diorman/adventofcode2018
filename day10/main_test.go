package main

import (
	"strings"
	"testing"
)

func TestShortestBoundaries(t *testing.T) {
	lights, err := parseInputFile("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected := 3
	if _, s := shortestBoundaries(lights); s != expected {
		t.Errorf("got: %v, wanted: %v", s, expected)
	}
}

func TestRender(t *testing.T) {
	lights, err := parseInputFile("input_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	b, s := shortestBoundaries(lights)
	expected := strings.Join([]string{
		"..................",
		"..................",
		"..................",
		"..................",
		"....#...#..###....",
		"....#...#...#.....",
		"....#...#...#.....",
		"....#####...#.....",
		"....#...#...#.....",
		"....#...#...#.....",
		"....#...#...#.....",
		"....#...#..###....",
		"..................",
		"..................",
		"..................",
		"..................",
	}, "\n")
	if f := render(lights, b, s, 4); f != expected {
		t.Errorf("got:\n%s\nwant:\n%s", f, expected)
	}
}
