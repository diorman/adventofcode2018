package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type boundaries struct {
	left, right int
	top, bottom int
}

func (b boundaries) area() int {
	return (b.right - b.left) * (b.bottom - b.top)
}

type xy struct {
	x, y int
}

type light struct {
	position, velocity xy
}

func (l light) positionAt(s int) xy {
	return xy{l.position.x + s*l.velocity.x, l.position.y + s*l.velocity.y}
}

func render(lights []light, b boundaries, sec, padding int) string {
	m := make(map[xy]int)
	for _, l := range lights {
		pos := l.positionAt(sec)
		m[pos]++
	}
	var strb strings.Builder
	for y := b.top - padding; y <= b.bottom+padding; y++ {
		for x := b.left - padding; x <= b.right+padding; x++ {
			pos := xy{x, y}
			if m[pos] > 0 {
				strb.WriteRune('#')
			} else {
				strb.WriteRune('.')
			}
		}
		if y < b.bottom+padding {
			strb.WriteString("\n")
		}
	}
	return strb.String()
}

func boundariesAt(lights []light, sec int) boundaries {
	b := boundaries{}
	for i, l := range lights {
		pos := l.positionAt(sec)

		if i == 0 || pos.x < b.left {
			b.left = pos.x
		}
		if i == 0 || pos.x > b.right {
			b.right = pos.x
		}
		if i == 0 || pos.y < b.top {
			b.top = pos.y
		}
		if i == 0 || pos.y > b.bottom {
			b.bottom = pos.y
		}
	}
	return b
}

func shortestBoundaries(lights []light) (boundaries, int) {
	b := boundariesAt(lights, 0)
	for sec := 0; ; sec++ {
		tmp := boundariesAt(lights, sec+1)
		if tmp.area() > b.area() {
			return b, sec
		}
		b = tmp
	}
}

func parseInputFile(file string) ([]light, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var (
		scanner = bufio.NewScanner(f)
		result  = []light{}
	)
	for scanner.Scan() {
		var p, v xy
		_, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &p.x, &p.y, &v.x, &v.y)
		if err != nil {
			return nil, err
		}
		result = append(result, light{p, v})
	}
	return result, scanner.Err()
}

func main() {
	lights, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	b, s := shortestBoundaries(lights)
	fmt.Printf("After %d seconds:\n%s", s, render(lights, b, s, 4))
}
