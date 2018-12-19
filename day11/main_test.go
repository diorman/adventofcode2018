package main

import (
	"fmt"
	"testing"
)

func TestHologramPowerLevelAt(t *testing.T) {
	tt := []struct {
		coordinate   coordinate
		serialNumber int
		out          int
	}{
		{coordinate{3, 5}, 8, 4},
		{coordinate{122, 79}, 57, -5},
		{coordinate{217, 196}, 39, 0},
		{coordinate{101, 153}, 71, 4},
	}
	for _, tc := range tt {
		h := newHologram(tc.serialNumber)
		t.Run(fmt.Sprintf("%+v;%d", tc.coordinate, tc.serialNumber), func(t *testing.T) {
			if out := h.powerLevelAt(tc.coordinate); out != tc.out {
				t.Errorf("got: %v, expected: %v", out, tc.out)
			}
		})
	}
}

func TestHologramFindSquareCoordinate(t *testing.T) {
	tt := []struct {
		serialNumber int
		c            coordinate
		max          int
	}{
		{18, coordinate{33, 45}, 29},
		{42, coordinate{21, 61}, 30},
	}
	for _, tc := range tt {
		h := newHologram(tc.serialNumber)
		t.Run(fmt.Sprintf("%d", tc.serialNumber), func(t *testing.T) {
			if c, max := h.findSquareCoordinate(3); c != tc.c || max != tc.max {
				t.Errorf("got: {coordinate: %+v max: %v}, expected: {coordinate: %+v max: %v}", c, max, tc.c, tc.max)
			}
		})
	}
}

func TestHologramFindDynamicSquareCoordinate(t *testing.T) {
	tt := []struct {
		serialNumber int
		c            coordinate
		s            int
	}{
		{18, coordinate{90, 269}, 16},
		{42, coordinate{232, 251}, 12},
	}
	for _, tc := range tt {
		h := newHologram(tc.serialNumber)
		t.Run(fmt.Sprintf("%d", tc.serialNumber), func(t *testing.T) {
			if c, s := h.findDynamicSquareCoordinate(); c != tc.c || s != tc.s {
				t.Errorf("got: %v, expected: %v",
					[]int{c.x, c.y, s},
					[]int{tc.c.x, tc.c.y, tc.s})
			}
		})
	}
}
