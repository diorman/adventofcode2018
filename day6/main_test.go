package main

import (
	"testing"
)

var coordinates = []coordinate{
	{1, 1},
	{1, 6},
	{8, 3},
	{3, 4},
	{5, 5},
	{8, 9},
}

func TestGetLargestAreaSize(t *testing.T) {
	expected := 17
	if result := getLargestAreaSize(coordinates); result != expected {
		t.Errorf("got: %v, want: %v\n", result, expected)
	}
}

func TestGetRegionArea(t *testing.T) {
	expected := 16
	if result := getRegionArea(coordinates, 32); result != expected {
		t.Errorf("got: %v, want: %v\n", result, expected)
	}
}
