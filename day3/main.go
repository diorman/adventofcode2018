package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type claim struct {
	id     string
	top    int
	left   int
	width  int
	height int
}

func parseInputFile(file string) ([]claim, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		claims  []claim
		scanner = bufio.NewScanner(f)
	)
	for scanner.Scan() {
		c := claim{}
		fmt.Sscanf(scanner.Text(), "#%s @ %d,%d: %dx%d", &c.id, &c.left, &c.top, &c.width, &c.height)
		claims = append(claims, c)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return claims, nil
}

func countOverlaps(claims []claim) int {
	m := make(map[string]int)
	counter := 0
	for _, c := range claims {
		for i := c.left; i < c.left+c.width; i++ {
			for j := c.top; j < c.top+c.height; j++ {
				k := fmt.Sprintf("%d,%d", i, j)
				x, ok := m[k]
				if ok && x == 1 {
					counter++
					m[k] = -1
					continue
				}
				if !ok {
					m[k] = 1
				}
			}
		}
	}
	return counter
}

func findUniqueClaimID(claims []claim) string {
	m := make(map[string]string)
	claimTracker := make(map[string]bool)
	for _, c := range claims {
		claimTracker[c.id] = true
		for i := c.left; i < c.left+c.width; i++ {
			for j := c.top; j < c.top+c.height; j++ {
				k := fmt.Sprintf("%d,%d", i, j)
				if _, ok := m[k]; ok {
					delete(claimTracker, m[k])
					delete(claimTracker, c.id)
				}
				m[k] = c.id
			}
		}
	}
	for k := range claimTracker {
		return k
	}
	return ""
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	claims, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	switch *part {
	case "1":
		fmt.Println(countOverlaps(claims))
	case "2":
		fmt.Println(findUniqueClaimID(claims))
	}
}
