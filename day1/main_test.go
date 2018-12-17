package main

import (
	"fmt"
	"testing"
)

func TestCalculateFrecuency(t *testing.T) {
	tt := []struct {
		in  []int
		out int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			if out := calculateFrecuency(tc.in); out != tc.out {
				t.Errorf("got: %v, want: %v", out, tc.out)
			}
		})
	}
}

func TestFindRepeatedFrecuency(t *testing.T) {
	tt := []struct {
		in  []int
		out int
	}{
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			if out := findRepeatedFrecuency(tc.in); out != tc.out {
				t.Errorf("got: %v, want: %v", out, tc.out)
			}
		})
	}
}
