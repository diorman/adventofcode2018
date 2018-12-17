package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func maxKey(m map[int]int) int {
	var (
		key   int
		value int
	)
	for k, v := range m {
		if v > value {
			key = k
			value = v
		}
	}
	return key
}

type strategy func(minuteCount int, guardCount int) int

func strategy1(minuteCount int, guardCount int) int {
	return guardCount
}

func strategy2(minuteCount int, guardCount int) int {
	return minuteCount
}

func getGuardIDMultipliedByMinute(logs []string, s strategy) int {
	var (
		r              = regexp.MustCompile(`\d{2}:(\d{2})\]\s(.+)$`)
		minuteTracker  = make(map[int]map[int]int)
		guardTracker   = make(map[int]int)
		currentGuardID int
		sleepyGuard    = struct{ id, count int }{}
	)

	for i, l := range logs {
		ss := r.FindStringSubmatch(l)
		if strings.HasPrefix(ss[2], "Guard") {
			fmt.Sscanf(ss[2], "Guard #%d begins shift", &currentGuardID)
		}
		if ss[2] != "falls asleep" {
			continue
		}
		minute, _ := strconv.Atoi(ss[1])
		nextLogMinute := 60
		if i+1 < len(logs) {
			nextLogMinute, _ = strconv.Atoi(r.FindStringSubmatch(logs[i+1])[1])
		}

		for j := minute; j < nextLogMinute; j++ {
			if _, ok := minuteTracker[currentGuardID]; !ok {
				minuteTracker[currentGuardID] = make(map[int]int)
			}
			minuteTracker[currentGuardID][j]++
			guardTracker[currentGuardID]++
			count := s(minuteTracker[currentGuardID][j], guardTracker[currentGuardID])
			if count > sleepyGuard.count {
				sleepyGuard.id = currentGuardID
				sleepyGuard.count = count
			}
		}
	}

	return sleepyGuard.id * maxKey(minuteTracker[sleepyGuard.id])
}

func parseInputFile(file string) ([]string, error) {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(s)), "\n")
	sort.Strings(lines)
	return lines, nil
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	logs, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	switch *part {
	case "1":
		fmt.Println(getGuardIDMultipliedByMinute(logs, strategy1))
	case "2":
		fmt.Println(getGuardIDMultipliedByMinute(logs, strategy2))
	}
}
