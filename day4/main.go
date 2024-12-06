package main

import (
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func parse(input string) ([][]rune, []point) {
	var ret [][]rune
	var pos []point
	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var slice []rune
		for x, r := range line {
			slice = append(slice, r)
			if r == 'X' {
				pos = append(pos, point{x, y})
			}
		}
		ret = append(ret, slice)
	}
	return ret, pos
}

var dirs [][]int = [][]int{
	{-1, -1, -2, -2, -3, -3}, // nw
	{0, -1, 0, -2, 0, -3},    // n
	{+1, -1, +2, -2, +3, -3}, //ne
	{+1, 0, +2, 0, +3, 0},    // e
	{+1, +1, +2, +2, +3, +3}, // se
	{0, +1, 0, +2, 0, +3},    // s
	{-1, +1, -2, +2, -3, +3}, // sw
	{-1, 0, -2, 0, -3, 0},    // w
}

func xmas(p point, rs [][]rune) int {
	c := 0
	for _, dir := range dirs {
		s := "X"
		for ; len(dir) >= 2; dir = dir[2:] {
			nx := p.x + dir[0]
			ny := p.y + dir[1]
			if nx < 0 || nx > len(rs[0])-1 {
				break
			}
			if ny < 0 || ny > len(rs)-1 {
				break
			}
			s += string(rs[ny][nx])
		}
		if s == "XMAS" {
			c++
		}
	}
	return c
}

func silver(input string) int {
	s := 0
	runes, points := parse(input)
	for _, point := range points {
		s += xmas(point, runes)
	}
	return s
}

func mas(p point, rs [][]rune, as map[point]int) {
	for n, dir := range dirs {
		if n%2 == 1 {
			continue
		}
		s := "M"
		var a point
		for ; len(dir) > 2; dir = dir[2:] {
			nx := p.x + dir[0]
			ny := p.y + dir[1]
			if nx < 0 || nx > len(rs[0])-1 {
				break
			}
			if ny < 0 || ny > len(rs)-1 {
				break
			}
			if rs[ny][nx] == 'A' {
				a = point{x: nx, y: ny}
			}
			s += string(rs[ny][nx])
		}
		if s == "MAS" {
			as[a]++
		}
	}
}

func gold(input string) int {
	runes, _ := parse(input)
	var points = make(map[point]int)

	for y := range runes {
		for x := range runes[0] {
			if runes[y][x] == 'M' {
				mas(point{x, y}, runes, points)
			}
		}
	}
	s := 0
	for _, v := range points {
		if v > 1 {
			s++
		}
	}

	// fmt.Printf("%v\n", points)
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
	test := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	println(silver(test))
	println(gold(test))
	println(silver(read()))
	println(gold(read()))

}
