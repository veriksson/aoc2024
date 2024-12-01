package main

import (
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(input string) ([]int, []int) {
	var left []int
	var right []int

	data := strings.Fields(input)

	for ; len(data) > 1; data = data[2:] {
		lefti, err := strconv.Atoi(data[0])
		if err != nil {
			panic(err)
		}

		righti, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		left = append(left, lefti)
		right = append(right, righti)

	}
	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func silver(left, right []int) int {
	s := 0
	for i := 0; i < len(left); i++ {
		s += abs(left[i], right[i])
	}
	return s
}

func gold(left, right []int) int {
	s := 0
	lookup := make(map[int]int)
	for i := 0; i < len(right); i++ {
		lookup[right[i]]++
	}
	for i := 0; i < len(left); i++ {
		s += left[i] * lookup[left[i]]
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
	// testinput := `3   4
	// 4   3
	// 2   5
	// 1   3
	// 3   9
	// 3   3`

	left, right := parse(read())

	println(silver(left, right))

	println(gold(left, right))
}
