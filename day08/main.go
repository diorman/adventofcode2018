package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInputFile(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	result := []int{}

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return result, nil
}

func sum(ns ...int) int {
	sum := 0
	for _, n := range ns {
		sum += n
	}
	return sum
}

func sumNodes(input []int) int {
	sum, _ := sumTree(input, func(children []int, metadata []int) int {
		if len(children) == 0 {
			return sum(append(children, metadata...)...)
		}
		s := 0
		for _, i := range metadata {
			if len(children) >= i {
				s += children[i-1]
			}
		}
		return s
	})
	return sum
}

func sumMetadata(input []int) int {
	sum, _ := sumTree(input, func(children []int, metadata []int) int {
		return sum(append(children, metadata...)...)
	})
	return sum
}

func sumTree(input []int, fn func([]int, []int) int) (int, int) {
	var (
		nChildren, nMetadata = input[0], input[1]
		index                = 2
		children             = make([]int, nChildren)
		metadata             = make([]int, nMetadata)
	)

	for i := 0; i < nChildren; i++ {
		child, read := sumTree(input[index:], fn)
		children[i] = child
		index += read
	}

	for i := 0; i < nMetadata; i++ {
		metadata[i] = input[i+index]
	}

	return fn(children, metadata), index + nMetadata
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	input, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	switch *part {
	case "1":
		fmt.Println(sumMetadata(input))
	case "2":
		fmt.Println(sumNodes(input))
	}
}
