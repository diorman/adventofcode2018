package main

// Improved with https://en.wikipedia.org/wiki/Summed-area_table

import (
	"flag"
	"fmt"
	"sync"
)

type coordinate struct {
	x, y int
}

type hologram struct {
	size         int
	serialNumber int
	m            map[coordinate]int
}

func newHologram(serialNumber int) hologram {
	h := hologram{300, serialNumber, make(map[coordinate]int)}
	h.init()
	return h
}

func (h hologram) init() {
	for x := 1; x <= h.size; x++ {
		for y := 1; y <= h.size; y++ {
			c := coordinate{x, y}
			h.m[c] = h.powerLevelAt(c) +
				h.m[coordinate{x, y - 1}] +
				h.m[coordinate{x - 1, y}] -
				h.m[coordinate{x - 1, y - 1}]
		}
	}
}

func (h hologram) powerLevelAt(c coordinate) int {
	rackID := (c.x + 10)
	level := rackID * c.y
	level += h.serialNumber
	level *= rackID
	level = (level / 100) % 10
	level -= 5
	return level
}

func (h hologram) findSquareCoordinate(squareSize int) (coordinate, int) {
	var (
		c   coordinate
		max int
	)
	for x := 1; x <= h.size-squareSize; x++ {
		for y := 1; y <= h.size-squareSize; y++ {
			p := h.m[coordinate{x + squareSize, y + squareSize}] -
				h.m[coordinate{x, y + squareSize}] -
				h.m[coordinate{x + squareSize, y}] +
				h.m[coordinate{x, y}]
			if p > max {
				max, c = p, coordinate{x + 1, y + 1}
			}
		}
	}
	return c, max
}

func (h hologram) findDynamicSquareCoordinate() (coordinate, int) {
	var (
		xy   coordinate
		max  int
		size int
		wg   sync.WaitGroup
		mx   sync.Mutex
	)
	wg.Add(h.size)
	for s := 1; s <= h.size; s++ {
		go func(s int) {
			defer wg.Done()
			c, p := h.findSquareCoordinate(s)
			mx.Lock()
			defer mx.Unlock()
			if p > max {
				max, xy, size = p, c, s
			}
		}(s)
	}
	wg.Wait()
	return xy, size
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	h := newHologram(9435)

	switch *part {
	case "1":
		c, _ := h.findSquareCoordinate(3)
		fmt.Printf("%d,%d\n", c.x, c.y)
	case "2":
		c, n := h.findDynamicSquareCoordinate()
		fmt.Printf("%d,%d,%d\n", c.x, c.y, n)
	}
}
