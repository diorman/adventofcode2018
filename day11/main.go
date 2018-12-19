package main

import (
	"flag"
	"fmt"
	"math"
)

type coordinate struct {
	x, y int
}

type hologram struct {
	size         int
	serialNumber int
}

func newHologram(serialNumber int) hologram {
	return hologram{
		size:         300,
		serialNumber: serialNumber,
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

func (h hologram) fixedFuelSquarePower(c coordinate, squareSize int) int {
	sum := 0
	for x := c.x; x < c.x+squareSize; x++ {
		for y := c.y; y < c.y+squareSize; y++ {
			sum += h.powerLevelAt(coordinate{x, y})
		}
	}
	return sum
}

func (h hologram) dynamicFuelSquarePower(c coordinate) (int, int) {
	var sum, max, size int
	maxSquareSize := h.size - int(math.Max(float64(c.x), float64(c.y)))
	for d := 0; d < maxSquareSize; d++ {
		x, y := c.x, c.y+d
		for i := 0; i < d*2+1; i++ {
			sum += h.powerLevelAt(coordinate{x, y})
			if x < c.x+d {
				x++
			} else {
				y--
			}
		}

		if sum > max {
			max, size = sum, d+1
		}
	}
	return max, size
}

func (h hologram) findFixedFuelSquarePowerCoordinate(squareSize int) coordinate {
	var (
		result coordinate
		max    int
	)
	for x := 1; x <= h.size-squareSize; x++ {
		for y := 1; y <= h.size-squareSize; y++ {
			c := coordinate{x, y}
			if p := h.fixedFuelSquarePower(c, squareSize); p > max {
				max, result = p, c
			}
		}
	}
	return result
}

func (h hologram) findDynamicFuelSquarePowerCoordinate() (coordinate, int) {
	var (
		xy   coordinate
		max  int
		size int
	)
	for x := 1; x <= h.size; x++ {
		for y := 1; y <= h.size; y++ {
			c := coordinate{x, y}
			if p, s := h.dynamicFuelSquarePower(c); p > max {
				max, xy, size = p, c, s
			}
		}
	}
	return xy, size
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	h := newHologram(9435)

	switch *part {
	case "1":
		c := h.findFixedFuelSquarePowerCoordinate(3)
		fmt.Printf("%d,%d\n", c.x, c.y)
	case "2":
		c, n := h.findDynamicFuelSquarePowerCoordinate()
		fmt.Printf("%d,%d,%d\n", c.x, c.y, n)
	}
}
