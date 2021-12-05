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

type vents map[point]int

func (v vents) overlap(amount int) int {
	count := 0

	for _, value := range v {
		if value >= amount {
			count++
		}
	}

	return count
}

func (l line) points() []point {
	result := []point{}

	x := l.p1.x
	y := l.p1.y

	dx, dy := l.slope()

	for {
		pt := point{x: x, y: y}
		result = append(result, pt)

		x += dx
		y += dy

		if x == l.p2.x+dx && y == l.p2.y+dy {
			break
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

// lines will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees
func (l line) slope() (int, int) {
	var x, y int

	if l.p1.x < l.p2.x {
		x = 1
	} else if l.p1.x > l.p2.x {
		x = -1
	}

	if l.p1.y < l.p2.y {
		y = 1
	} else if l.p1.y > l.p2.y {
		y = -1
	}

	return x, y
}

func toi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func numeric(c rune) bool {
	return !unicode.IsNumber(c)
}

func parseLines(input []string) []line {
	result := []line{}

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

		result = append(result, l)
	}

	return result
}

func SolvePart1(input []string) int {
	lines := parseLines(input)

	vents := vents{}
	for _, line := range lines {
		if line.vertical() || line.horizontal() {
			for _, point := range line.points() {
				vents[point] += 1
			}
		}
	}

	return vents.overlap(2)
}

func SolvePart2(input []string) int {
	lines := parseLines(input)

	vents := vents{}
	for _, line := range lines {
		for _, point := range line.points() {
			vents[point] += 1
		}
	}

	return vents.overlap(2)
}
