package main

import (
	"fmt"
	"strings"
)

func parseNotes(notes []string) map[string]bool {
	m := make(map[string]bool)
	for _, n := range notes {
		if n[9] == byte('#') {
			m[n[0:5]] = true
		}
	}
	return m
}

func doGeneration(state string, initPos int, notes map[string]bool) (string, int) {
	state = "...." + state[:strings.LastIndex(state, "#")+1] + "...."
	var builder strings.Builder
	var firstIndex int
	for j := 0; j < len(state)-5; j++ {
		if notes[state[j:j+5]] {
			if builder.Len() == 0 {
				firstIndex = j
			}
			builder.WriteRune('#')
			continue
		}
		if builder.Len() > 0 {
			builder.WriteByte('.')
		}
	}
	return builder.String(), initPos + firstIndex - 2
}

func sumPotNumbers(initialState string, notes []string, nGenerations int) int {
	var (
		pnotes  = parseNotes(notes)
		state   = initialState
		initPos = 0
	)

	for i := 0; i < nGenerations; i++ {
		state, initPos = doGeneration(state, initPos, pnotes)
	}

	sum := 0
	for n, char := range state {
		if char == '#' {
			sum += n + initPos
		}
	}
	return sum
}

func main() {
	initialState := "##...#......##......#.####.##.#..#..####.#.######.##..#.####...##....#.#.####.####.#..#.######.##..."
	notes := []string{
		"#.... => .",
		"#..## => #",
		"....# => .",
		"...#. => .",
		"...## => #",
		"#.#.# => .",
		".#... => #",
		"##.#. => .",
		"..#.# => .",
		".##.# => #",
		"###.# => #",
		".#.## => .",
		"..... => .",
		"##### => #",
		"###.. => .",
		"##..# => #",
		"#.### => #",
		"#.#.. => .",
		"..### => .",
		"..#.. => .",
		".#..# => #",
		".##.. => #",
		"##... => #",
		".#.#. => #",
		".###. => #",
		"#..#. => .",
		"####. => .",
		".#### => #",
		"#.##. => #",
		"##.## => .",
		"..##. => .",
		"#...# => #",
	}

	var (
		sum20          = sumPotNumbers(initialState, notes, 20)
		sum100         = sumPotNumbers(initialState, notes, 100)
		sum101         = sumPotNumbers(initialState, notes, 101)
		sum50000000000 = (50000000000-100)*(sum101-sum100) + sum100
	)
	fmt.Printf("part 1: %d, part 2: %d\n", sum20, sum50000000000)
}
