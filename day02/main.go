package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkLetters(id string) (hasTwoLetters, hasThreeLetters bool) {
	counts := make(map[rune]int)

	for _, r := range id {
		counts[r]++
	}

	for _, c := range counts {
		switch c {
		case 2:
			hasTwoLetters = true
		case 3:
			hasThreeLetters = true
		}
		if hasTwoLetters && hasThreeLetters {
			break
		}
	}

	return
}

func calculateChecksum(ids []string) int {
	var (
		twoLetterCount   = 0
		threeLetterCount = 0
	)

	for _, id := range ids {
		hasTwoLetters, hasThreeLetters := checkLetters(id)

		if hasTwoLetters {
			twoLetterCount++
		}

		if hasThreeLetters {
			threeLetterCount++
		}
	}

	return twoLetterCount * threeLetterCount
}

func parseInputFile(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		result  = []string{}
		scanner = bufio.NewScanner(f)
	)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

type node struct {
	value    string
	children map[rune]*node
}

type trie struct {
	root *node
}

func (t trie) add(value string) {
	var (
		n       = t.root
		newNode *node
	)
	for _, c := range value {
		if _, ok := n.children[c]; ok {
			n = n.children[c]
			continue
		}
		newNode = &node{children: make(map[rune]*node)}
		n.children[c] = newNode
		n = newNode
	}
	newNode.value = value
}

func (t trie) findClose(value string) (string, bool) {
	return findClose(t.root, value, false)
}

func findClose(n *node, value string, diffFound bool) (string, bool) {
	r := rune(value[0])

	for cr, child := range n.children {
		if len(value) == 1 && cr == r {
			return child.value, true
		}
		if len(value) == 1 || (cr != r && diffFound) {
			continue
		}
		if v, ok := findClose(child, value[1:], diffFound || cr != r); ok {
			return v, true
		}
	}
	return "", false
}

func getCommonLetters(ids []string) string {
	var (
		t        = trie{&node{children: make(map[rune]*node)}}
		ida, idb string
		strb     strings.Builder
	)

	for _, id := range ids {
		var ok bool
		if ida, ok = t.findClose(id); ok {
			idb = id
			break
		}
		t.add(id)
	}

	for i := 0; i < len(ida); i++ {
		if ida[i] == idb[i] {
			strb.WriteByte(ida[i])
		}
	}

	return strb.String()
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	ids, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	switch *part {
	case "1":
		fmt.Println(calculateChecksum(ids))
	case "2":
		fmt.Println(getCommonLetters(ids))
	}
}
