package main

import (
	"aoc/internal/data"
	"fmt"
	"strconv"
	"strings"
)

const size = 5

type board struct {
	numbers [size][size]int
	marks   [size][size]bool
	won     bool
}

func (b *board) mark(v int) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if b.numbers[row][col] == v {
				b.marks[row][col] = true
				return
			}
		}
	}
}

func (b *board) win() bool {
	if b.won {
		return true
	}

	// rows
	for row := 0; row < size; row++ {
		marked := 0
		for col := 0; col < size; col++ {
			if b.marks[row][col] {
				marked++
			}

			if marked == 5 {
				b.won = true
				return true
			}
		}
	}

	// columns
	for col := 0; col < size; col++ {
		marked := 0
		for row := 0; row < size; row++ {
			if b.marks[row][col] {
				marked++
			}

			if marked == 5 {
				b.won = true
				return true
			}
		}
	}

	return false
}

func (b board) score() int {
	score := 0

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if !b.marks[row][col] {
				score += b.numbers[row][col]
			}
		}
	}

	return score
}

func main() {
	input := data.LoadStrings()

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

func split(input string, sep string) []int {
	result := []int{}
	var parts []string

	if sep == "" {
		parts = strings.Fields(input)
	} else {
		parts = strings.Split(input, sep)
	}

	for _, str := range parts {
		n, _ := strconv.Atoi(str)
		result = append(result, n)
	}

	return result
}

func parsebingo(input []string) ([]int, []board) {
	// parse selections
	selections := split(input[0], ",")

	// parse boards
	boards := []board{}
	currentBoard := board{}
	i := 0
	for _, line := range input[2:] {
		if line == "" {
			boards = append(boards, currentBoard)
			currentBoard = board{}
			i = 0
			continue
		}

		nums := split(line, "")
		for j, val := range nums {
			currentBoard.numbers[i][j] = val
		}

		i++
	}

	return selections, boards
}

func allwon(boards []board) bool {
	for i := range boards {
		if !boards[i].win() {
			return false
		}
	}

	return true
}

func SolvePart1(input []string) int {
	selections, boards := parsebingo(input)

	for _, s := range selections {
		for i := range boards {
			boards[i].mark(s)

			if boards[i].win() {
				return boards[i].score() * s
			}
		}
	}

	return 0
}

func SolvePart2(input []string) int {
	selections, boards := parsebingo(input)

	for _, s := range selections {
		for i := range boards {
			boards[i].mark(s)

			if allwon(boards) {
				return boards[i].score() * s
			}
		}
	}

	return 0
}
