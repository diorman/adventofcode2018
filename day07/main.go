package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type step struct {
	id       rune
	parents  steps
	children steps
	complete bool
}

type steps []*step

type worker struct {
	step    *step
	seconds int
	busy    bool
}

func (ss steps) areComplete() bool {
	for _, n := range ss {
		if !n.complete {
			return false
		}
	}
	return true
}

func (ss steps) getAvailable() steps {
	r := steps{}
	for _, s := range ss {
		if s.complete || !s.parents.areComplete() {
			continue
		}
		r = append(r, s)
	}
	return r
}

func (ss steps) getNextIndex() int {
	var index = -1
	for i, s := range ss {
		if index == -1 || s.id < ss[index].id {
			index = i
		}
	}
	return index
}

func (ss steps) getNext() *step {
	i := ss.getNextIndex()
	if i == -1 {
		return nil
	}
	return ss[i]
}

func (ss steps) removeAt(i int) steps {
	return append(ss[:i], ss[i+1:]...)
}

func (ss steps) remove(step *step) steps {
	for i, s := range ss {
		if s == step {
			return ss.removeAt(i)
		}
	}
	return ss
}

func (ss steps) sort() steps {
	available := ss.getAvailable()
	ordered := steps{}

	for len(available) > 0 {
		i := available.getNextIndex()
		s := available[i]
		s.complete = true
		ordered = append(ordered, s)
		available = append(available.removeAt(i), s.children.getAvailable()...)
	}

	return ordered
}

func (ss steps) duration(nworkers int, baseTime int) int {
	var (
		available   = ss.getAvailable()
		seconds     = 0
		busyWorkers = 0
		workers     = []*worker{}
	)

	for i := 0; i < nworkers; i++ {
		workers = append(workers, &worker{})
	}

	for len(available) > 0 || busyWorkers > 0 {
		done := steps{}

		for _, w := range workers {
			if !w.busy {
				s := available.getNext()
				if s == nil {
					continue
				}
				w.step, w.seconds, w.busy = s, 0, true
				available = append(available.remove(s))
				busyWorkers++
			}
			if w.seconds == (int(w.step.id)-65)+baseTime {
				done = append(done, w.step)
				w.busy = false
				busyWorkers--
			}
			w.seconds++
		}
		for _, s := range done {
			s.complete = true
			available = append(available, s.children.getAvailable()...)
		}
		seconds++
	}

	return seconds
}

func (ss steps) String() string {
	var ids strings.Builder
	for _, s := range ss {
		ids.WriteRune(s.id)
	}
	return ids.String()
}

func parseInputFile(file string) (steps, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var (
		scanner = bufio.NewScanner(f)
		m       = make(map[rune]*step)
	)

	for scanner.Scan() {
		var parent, child rune
		_, err := fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.", &parent, &child)
		if err != nil {
			return nil, err
		}
		if _, ok := m[parent]; !ok {
			m[parent] = &step{id: parent}
		}
		if _, ok := m[child]; !ok {
			m[child] = &step{id: child}
		}
		m[child].parents = append(m[child].parents, m[parent])
		m[parent].children = append(m[parent].children, m[child])
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	steps := steps{}
	for _, n := range m {
		steps = append(steps, n)
	}
	return steps, nil
}

func main() {
	part := flag.String("part", "1", "")
	flag.Parse()

	steps, err := parseInputFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	switch *part {
	case "1":
		fmt.Println(steps.sort())
	case "2":
		fmt.Println(steps.duration(5, 60))
	}
}
