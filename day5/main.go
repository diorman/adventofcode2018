package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
)

type node struct {
	prev  *node
	value rune
}

func getPolimerLength(data string) int {
	deleted := 0
	var n *node
	for _, r := range data {
		if n == nil || math.Abs(float64(n.value-r))-32 != 0 {
			n = &node{n, r}
			continue
		}
		deleted += 2
		n = n.prev
	}

	return len(data) - deleted
}

func getImprovedPolimerLength(data string) int {
	min := 0
	for r := 'a'; r <= 'z'; r++ {
		r := regexp.MustCompile(fmt.Sprintf("(?i)%c", r))
		d := r.ReplaceAllString(data, "")
		l := getPolimerLength(d)
		if min == 0 || l < min {
			min = l
		}
	}
	return min
}

func parseInputFile(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	return string(bytes), err
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	data, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	switch *part {
	case "1":
		fmt.Println(getPolimerLength(data))
	case "2":
		fmt.Println(getImprovedPolimerLength(data))
	}
}
