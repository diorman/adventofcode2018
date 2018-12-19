package main

import (
	"fmt"
	"testing"
)

func TestGetHighestScore(t *testing.T) {
	tt := []struct {
		nPlayers int
		nMarbles int
		maxScore int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d,%d", tc.nPlayers, tc.nMarbles), func(t *testing.T) {
			if r := getHighestScore(tc.nPlayers, tc.nMarbles); r != tc.maxScore {
				t.Errorf("got: %v, wanted: %v", r, tc.maxScore)
			}
		})
	}
}
