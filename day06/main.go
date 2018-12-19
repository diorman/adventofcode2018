package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

type coordinate struct{ x, y int }

type boundaries struct{ left, right, top, bottom int }

type distance struct {
	coordinate coordinate
	length     int
}

type cellInfo struct {
	distances []distance
	isBorder  bool
}

func calculateDistance(a, b coordinate) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func isBorder(c coordinate, b boundaries) bool {
	return c.x == b.left || c.x == b.right || c.y == b.top || c.y == b.bottom
}

func getShortestDistance(distances []distance) (distance, bool) {
	var (
		m   = make(map[int][]distance)
		min = -1
	)

	for _, dist := range distances {
		m[dist.length] = append(m[dist.length], dist)
		if min == -1 || dist.length < min {
			min = dist.length
		}
	}

	if len(m[min]) == 1 {
		return m[min][0], true
	}

	return distance{}, false
}

func getBoundaries(coordinates []coordinate) boundaries {
	b := boundaries{}
	for i, c := range coordinates {
		if i == 0 {
			b.left = c.x
			b.right = c.x
			b.top = c.y
			b.bottom = c.y
		}
		if c.x < b.left {
			b.left = c.x
		}
		if c.x > b.right {
			b.right = c.x
		}
		if c.y < b.top {
			b.top = c.y
		}
		if c.y > b.bottom {
			b.bottom = c.y
		}
	}
	return b
}

func populateMap(coordinates []coordinate) map[coordinate]cellInfo {
	var (
		m = make(map[coordinate]cellInfo)
		b = getBoundaries(coordinates)
	)

	for x := b.left; x <= b.right; x++ {
		for y := b.top; y <= b.bottom; y++ {
			xy := coordinate{x, y}
			distances := []distance{}
			for _, c := range coordinates {
				dist := distance{c, calculateDistance(c, xy)}
				distances = append(distances, dist)
			}
			m[xy] = cellInfo{distances, isBorder(xy, b)}
		}
	}
	return m
}

func getLargestAreaSize(coordinates []coordinate) int {
	var (
		counter  = make(map[coordinate]int)
		max      int
		infinite = make(map[coordinate]bool)
		m        = populateMap(coordinates)
	)

	for _, ci := range m {
		d, ok := getShortestDistance(ci.distances)
		if !ok {
			continue
		}
		if ci.isBorder {
			infinite[d.coordinate] = true
		}
		counter[d.coordinate]++
	}

	for c, count := range counter {
		if _, ok := infinite[c]; ok {
			continue
		}
		if count > max {
			max = count
		}
	}
	return max
}

func getDistancesSum(distances []distance) int {
	sum := 0
	for _, dist := range distances {
		sum += dist.length
	}
	return sum
}

func getRegionArea(coordinates []coordinate, dist int) int {
	var (
		counter = 0
		m       = populateMap(coordinates)
	)
	for _, ci := range m {
		if getDistancesSum(ci.distances) < dist {
			counter++
		}
	}
	return counter
}

func parseInputFile(file string) ([]coordinate, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var (
		coordinates = []coordinate{}
		scanner     = bufio.NewScanner(f)
	)

	for scanner.Scan() {
		c := coordinate{}
		_, err := fmt.Sscanf(scanner.Text(), "%d, %d", &c.x, &c.y)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, c)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return coordinates, nil
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	coordinates, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalf("could parse input file: %v\n", err)
	}

	switch *part {
	case "1":
		fmt.Println(getLargestAreaSize(coordinates))
	case "2":
		fmt.Println(getRegionArea(coordinates, 10000))
	}
}
