package main

import (
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type printqueue struct {
	// a reverse rule list.
	// since we traverse the list going forward,
	// if we get a hit in the rules it means that we
	// encounter a number later that we should be after
	rules map[int]map[int]struct{}

	pages [][]int
}

func parse(input string) printqueue {
	l := printqueue{
		rules: make(map[int]map[int]struct{}),
	}

	lines := strings.Split(input, "\n")
	i := 0
	for ; len(lines[i]) > 0; i++ {
		parts := strings.Split(lines[i], "|")
		fst, _ := strconv.Atoi(parts[0])
		snd, _ := strconv.Atoi(parts[1])

		if m, ok := l.rules[snd]; ok {
			m[fst] = struct{}{}
		} else {
			l.rules[snd] = make(map[int]struct{})
			l.rules[snd][fst] = struct{}{}
		}
	}
	i++
	for ; i < len(lines); i++ {
		parts := strings.Split(lines[i], ",")
		var r []int
		for _, part := range parts {
			j, _ := strconv.Atoi(strings.TrimSpace(part))
			r = append(r, j)
		}
		l.pages = append(l.pages, r)
	}
	return l
}

func silver1(in printqueue) int {
	s := 0
	for _, p := range in.pages {
		valid := true
		for i := 0; i < len(p); i++ {
			if rules, ok := in.rules[p[i]]; ok {
				for j := i + 1; j < len(p); j++ {
					v := p[j]

					if _, ok := rules[v]; ok {
						// we found a value that should precede us
						valid = false
					}
				}
			}
		}
		if valid {
			m := len(p) / 2
			s += p[m]
		}
	}
	return s
}

// reducerules creates a reduced set of rules just for the
// pagenumbers that we are working with, since others are not
// needed to order the pages
func reducerules(ps []int, pq printqueue) map[int]int {
	rules := make(map[int]int)
	for _, n := range ps {
		rules[n] = 0
		rs := pq.rules[n]
		for _, m := range ps {
			if _, ok := rs[m]; ok {
				rules[n]++
			}
		}
	}
	return rules
}

// silver2 was made after I realized the solution for gold
func silver2(in printqueue) int {
	s := 0
	for _, p := range in.pages {
		cp := make([]int, len(p))
		copy(cp, p)

		rules := reducerules(p, in)
		sort.Slice(cp, func(i, j int) bool {
			v1 := rules[cp[i]]
			v2 := rules[cp[j]]
			return v1 < v2
		})

		// if there is no change to the array
		// it was in the correct order
		if reflect.DeepEqual(p, cp) {
			m := len(cp) / 2
			s += cp[m]
		}
	}
	return s
}

func gold(pq printqueue) int {
	s := 0
	for _, ps := range pq.pages {
		cp := make([]int, len(ps))
		copy(cp, ps)
		rules := reducerules(ps, pq)
		sort.Slice(cp, func(i, j int) bool {
			v1 := rules[cp[i]]
			v2 := rules[cp[j]]
			return v1 < v2
		})

		// if there is no change to the array
		// it was in the correct order
		if reflect.DeepEqual(ps, cp) {
			continue
		}

		m := len(cp) / 2
		s += cp[m]
	}
	return s
}

func read() string {
	d, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	return string(d)
}

func main() {
	test := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	parsedtest := parse(test)
	println(silver1(parsedtest))
	println(silver2(parsedtest))
	println(gold(parsedtest))

	parsed := parse(read())
	println(silver1(parsed))
	println(silver2(parsed))
	println(gold(parsed))
}
