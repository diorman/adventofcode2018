package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type coordinate struct {
	x, y int
}

type cartTrack struct {
	paths map[coordinate]*path
	carts []*cart
}

func (t *cartTrack) Len() int {
	return len(t.carts)
}

func (t *cartTrack) Less(i, j int) bool {
	ci, cj := t.carts[i], t.carts[j]
	if ci.position.y < cj.position.y {
		return true
	}
	if ci.position.y == cj.position.y && ci.position.x < cj.position.x {
		return true
	}
	return false
}

func (t *cartTrack) Swap(i, j int) {
	tmp := t.carts[i]
	t.carts[i], t.carts[j] = t.carts[j], tmp
}

func (t *cartTrack) tick() []coordinate {
	sort.Sort(t)

	crashLocations := []coordinate{}
	for _, c := range t.carts {
		if c.crashed {
			continue
		}
		t.paths[c.position].cart = nil
		c.move(*t.paths[c.position])
		if t.paths[c.position].cart != nil {
			c.crashed = true
			t.paths[c.position].cart.crashed = true
			crashLocations = append(crashLocations, c.position)
			t.paths[c.position].cart = nil
			continue
		}
		t.paths[c.position].cart = c
	}

	carts := []*cart{}
	for _, c := range t.carts {
		if !c.crashed {
			carts = append(carts, c)
		}
	}
	t.carts = carts
	return crashLocations
}

func (t *cartTrack) findFirstCrash() coordinate {
	for {
		if crashLocations := t.tick(); len(crashLocations) > 0 {
			return crashLocations[0]
		}
	}
}

func (t *cartTrack) findLastCartLocation() coordinate {
	for len(t.carts) > 1 {
		t.tick()
	}
	return t.carts[0].position
}

type cart struct {
	position     coordinate
	direction    rune
	intersection int
	crashed      bool
}

func (c *cart) move(p path) {
	c.updateDirection(p)
	c.updatePosition()
	if p.pathType == '+' {
		c.intersection++
	}
}

func (c *cart) updatePosition() {
	switch c.direction {
	case '^':
		c.position.y--
	case 'v':
		c.position.y++
	case '<':
		c.position.x--
	case '>':
		c.position.x++
	}
}

func (c *cart) turnRight() {
	directions := "^>v<^"
	c.direction = rune(directions[strings.IndexRune(directions, c.direction)+1])
}

func (c *cart) turnLeft() {
	directions := "^<v>^"
	c.direction = rune(directions[strings.IndexRune(directions, c.direction)+1])
}

func (c *cart) updateDirection(p path) {
	var (
		vertical        = strings.IndexRune("^v", c.direction) >= 0
		horizontal      = strings.IndexRune("<>", c.direction) >= 0
		intersectionMod = c.intersection % 3
	)
	if (p.pathType == '/' && horizontal) ||
		(p.pathType == '\\' && vertical) ||
		(p.pathType == '+' && intersectionMod == 0) {
		c.turnLeft()
		return
	}
	if (p.pathType == '/' && vertical) ||
		(p.pathType == '\\' && horizontal) ||
		(p.pathType == '+' && intersectionMod == 2) {
		c.turnRight()
		return
	}
}

type path struct {
	pathType rune
	cart     *cart
}

func newPath(r rune) *path {
	var (
		pathType = r
	)
	if r == '^' || r == 'v' {
		pathType = '|'
	}
	if r == '<' || r == '>' {
		pathType = '-'
	}
	return &path{pathType: pathType}
}

func parseInputFile(file string) (*cartTrack, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var (
		carts   = []*cart{}
		paths   = make(map[coordinate]*path)
		scanner = bufio.NewScanner(f)
		x, y    int
	)

	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			x = 0
			y++
			continue
		}
		var (
			position = coordinate{x, y}
			r        = rune(scanner.Text()[0])
			path     = newPath(r)
		)
		paths[position] = path
		if strings.IndexRune("<>^v", r) >= 0 {
			cart := &cart{position: position, direction: r}
			carts = append(carts, cart)
			path.cart = cart
		}
		x++
	}
	return &cartTrack{paths, carts}, scanner.Err()
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	t, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	switch *part {
	case "1":
		position := t.findFirstCrash()
		fmt.Printf("%d,%d", position.x, position.y)
	case "2":
		position := t.findLastCartLocation()
		fmt.Printf("%d,%d", position.x, position.y)
	}
}
