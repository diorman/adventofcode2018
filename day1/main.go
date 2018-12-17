package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFrecuency(input []int) int {
	frecuency := 0
	for _, n := range input {
		frecuency += n
	}
	return frecuency
}

func findRepeatedFrecuency(input []int) int {
	var (
		frecuencies = map[int]bool{0: true}
		frecuency   = 0
	)

	for {
		for _, n := range input {
			frecuency += n
			if _, exist := frecuencies[frecuency]; exist {
				return frecuency
			}
			frecuencies[frecuency] = true
		}
	}
}

func parseInputFile(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		result  = []int{}
		scanner = bufio.NewScanner(f)
	)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, scanner.Err()
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
		fmt.Println(calculateFrecuency(input))
	case "2":
		fmt.Println(findRepeatedFrecuency(input))
	}
}
