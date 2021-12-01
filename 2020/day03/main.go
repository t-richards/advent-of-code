package main

import (
	"bufio"
	"fmt"
	"os"
)

var partTwoSlopes = [][2]int{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(TreeCount(input, 3, 1))
	fmt.Println(TreeCountPartTwo(input))
}

func TreeCount(input []string, slopeX int, slopeY int) int {
	count := 0
	row := 0
	col := 0

	for row < len(input) {
		if input[row][col%len(input[row])] == '#' {
			count++
		}

		row += slopeY
		col += slopeX
	}

	return count
}

func TreeCountPartTwo(input []string) int {
	result := 1

	for _, slope := range partTwoSlopes {
		result *= TreeCount(input, slope[0], slope[1])
	}

	return result
}
