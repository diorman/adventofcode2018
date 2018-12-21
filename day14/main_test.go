package main

import (
	"strconv"
	"testing"
)

func TestLastTenExtraScores(t *testing.T) {
	tt := []struct {
		in  int
		out string
	}{
		{9, "5158916779"},
		{5, "0124515891"},
		{18, "9251071085"},
		{2018, "5941429882"},
	}
	for _, tc := range tt {
		t.Run(strconv.Itoa(tc.in), func(t *testing.T) {
			if out := lastTenExtraScores(tc.in); out != tc.out {
				t.Errorf("got: %v, expected: %v", out, tc.out)
			}
		})
	}
}

func TestLengthBefore(t *testing.T) {
	tt := []struct {
		in  string
		out int
	}{
		{"51589", 9},
		{"01245", 5},
		{"92510", 18},
		{"59414", 2018},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			if out := lengthBeforeKey(tc.in); out != tc.out {
				t.Errorf("got: %v, expected: %v", out, tc.out)
			}
		})
	}
}
