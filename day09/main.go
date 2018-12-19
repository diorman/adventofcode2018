package main

import (
	"container/ring"
	"flag"
	"fmt"
)

func newNode(value int) *ring.Ring {
	r := ring.New(1)
	r.Value = value
	return r
}

func getHighestScore(nPlayers, nMarbels int) int {
	var (
		max     = 0
		scores  = make(map[int]int)
		current = newNode(0)
	)

	for marble := 1; marble <= nMarbels; marble++ {
		playerID := marble % nPlayers
		if marble%23 == 0 {
			extra := current.Move(-7)
			current = extra.Next()
			extra.Prev().Unlink(1)

			scores[playerID] += marble + extra.Value.(int)
			if scores[playerID] > max {
				max = scores[playerID]
			}
		} else {
			n := newNode(marble)
			current.Next().Link(n)
			current = n
		}
	}
	return max
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	switch *part {
	case "1":
		fmt.Println(getHighestScore(459, 71320))
	case "2":
		fmt.Println(getHighestScore(459, 7132000))
	}
}
