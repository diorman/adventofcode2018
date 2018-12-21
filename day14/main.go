package main

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"
)

func moveElf(scoreboard []byte, elf int) int {
	return (1 + elf + int(scoreboard[elf]-48)) % len(scoreboard)
}

func combineRecipes(scoreboard []byte, elf1, elf2 int) []byte {
	new := strconv.Itoa(int(scoreboard[elf1]+scoreboard[elf2]) - 96)
	return append(scoreboard, new...)
}

func run(stop func([]byte) bool) []byte {
	scoreboard := []byte("37")
	elf1 := 0
	elf2 := 1
	for {
		scoreboard = combineRecipes(scoreboard, elf1, elf2)
		elf1 = moveElf(scoreboard, elf1)
		elf2 = moveElf(scoreboard, elf2)
		if stop(scoreboard) {
			break
		}
	}
	return scoreboard
}

func lengthBeforeKey(key string) int {
	index := 0
	keyb := []byte(key)
	run(func(scoreboard []byte) bool {
		if len(scoreboard)%1000000 == 0 {
			if index = bytes.Index(scoreboard, keyb); index >= 0 {
				return true
			}
		}
		return false
	})
	return index
}

func lastTenExtraScores(n int) string {
	scoreboard := run(func(scoreboard []byte) bool {
		return len(scoreboard) >= n+10
	})
	return string(scoreboard[n : n+10])
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	switch *part {
	case "1":
		fmt.Println(lastTenExtraScores(824501))
	case "2":
		fmt.Println(lengthBeforeKey("824501"))
	}
}
