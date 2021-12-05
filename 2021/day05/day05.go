package main

import (
	"aoc/internal/data"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := data.LoadStrings()

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

func (l line) points() []point {
	result := []point{}

	xmin, xmax := minmax(l.p1.x, l.p2.x)
	ymin, ymax := minmax(l.p1.y, l.p2.y)

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			pt := point{x: x, y: y}
			result = append(result, pt)
		}
	}

	return result
}

func (l line) vertical() bool {
	return l.p1.x == l.p2.x
}

func (l line) horizontal() bool {
	return l.p1.y == l.p2.y
}

func minmax(x, y int) (int, int) {
	if x < y {
		return x, y
	}

	return y, x
}

func toi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func numeric(c rune) bool {
	return !unicode.IsNumber(c)
}

func SolvePart1(input []string) int {
	lines := []line{}
	for _, s := range input {
		parts := strings.FieldsFunc(s, numeric)

		l := line{
			p1: point{
				x: toi(parts[0]),
				y: toi(parts[1]),
			},
			p2: point{
				x: toi(parts[2]),
				y: toi(parts[3]),
			},
		}

		// only consider horizontal and vertical lines
		if l.vertical() || l.horizontal() {
			lines = append(lines, l)
		}
	}

	vents := map[point]int{}
	for _, line := range lines {
		for _, point := range line.points() {
			vents[point] += 1
		}
	}

	count := 0
	for _, v := range vents {
		if v > 1 {
			count++
		}
	}

	return count
}

func SolvePart2(input []string) int {
	return 0
}
