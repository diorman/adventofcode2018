package main

import "testing"

func TestSumPotNumbers(t *testing.T) {
	initialState := "#..#.#..##......###...###"
	notes := []string{
		"...## => #",
		"..#.. => #",
		".#... => #",
		".#.#. => #",
		".#.## => #",
		".##.. => #",
		".#### => #",
		"#.#.# => #",
		"#.### => #",
		"##.#. => #",
		"##.## => #",
		"###.. => #",
		"###.# => #",
		"####. => #",
	}
	expected := 325
	if s := sumPotNumbers(initialState, notes, 20); s != expected {
		t.Errorf("got: %v, expected: %v", s, expected)
	}
}
