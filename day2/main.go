package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	var result [][]int
	for _, line := range lines {
		var ints []int
		for _, c := range strings.Fields(line) {
			i, _ := strconv.Atoi(c)
			ints = append(ints, i)
		}
		result = append(result, ints)
	}
	return result
}

func silver(input [][]int) int {
	s := 0

	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		if safe(line) {
			s++
		}
	}
	return s
}

func safe(ints []int) bool {
	asc := true
	if ints[0] < ints[1] {
		asc = false
	}

	for i := 1; i < len(ints); i++ {
		if asc && (ints[i-1] < ints[i]) {
			return false // breaking ascending
		}
		if !asc && (ints[i-1] > ints[i]) {
			return false
		}

		v := abs(ints[i-1], ints[i])
		if v < 1 || v > 3 {
			return false
		}
	}
	return true
}

func gold(input [][]int) int {
	s := 0
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		if safe(line) {
			s++
		} else {
			for i := range line {
				var t []int
				for j := 0; j < len(line); j++ {
					if j == i {
						continue
					}
					t = append(t, line[j])
				}
				if safe(t) {
					s++
					break
				}
			}
		}
	}
	return s
}

func abs(left, right int) int {
	v := math.Abs(float64(left) - float64(right))
	return int(v)
}

func read() string {
	d, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	return string(d)
}

func main() {
	// test := `7 6 4 2 1
	// 1 2 7 8 9
	// 9 7 6 2 1
	// 1 3 2 4 5
	// 8 6 4 4 1
	// 1 3 6 7 9`

	parsed := parse(read())
	// fmt.Printf("%v\n", parsed)
	println(silver(parsed))
	println(gold(parsed))

	// parsed2 := parse(test)
	// println(gold(parsed2))
}
