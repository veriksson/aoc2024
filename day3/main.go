package main

import (
	"os"
	"strconv"
	"strings"
)

func read() string {
	d, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	return string(d)
}

type state struct {
	s int
	e bool
}

type action func(*state)

func parse(input string) []action {
	var ret []action
	for i := 0; i < len(input)-7; i++ {
		if input[i:i+7] == "don't()" {
			ret = append(ret, func(s *state) {
				s.e = false
			})
			continue
		}
		if input[i:i+4] == "do()" {
			ret = append(ret, func(s *state) {
				s.e = true
			})
			continue
		}
		if input[i:i+4] == "mul(" {
			end := strings.IndexByte(input[i+5:], ')')
			if end == -1 || end >= 7 {
				continue
			}
			numbers := strings.Split(input[i+4:i+5+end], ",")
			if len(numbers) == 1 {
				continue
			}
			left, _ := strconv.Atoi(numbers[0])
			right, _ := strconv.Atoi(numbers[1])
			ret = append(ret, func(s *state) {
				if s.e {
					s.s += left * right
				}
			})
			i = i + end
		}
	}
	return ret
}

func silver(input string) int {
	s := &state{s: 0, e: true}
	actions := parse(input)
	for _, act := range actions {
		act(s)
	}
	return s.s
}

func main() {
	// test := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	// println(silver(test))
	println(silver("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
	println(silver(read()))
}
