package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(TreeCount1(input))
}

func TreeCount1(input []string) int {
	count := 0
	row := 0
	col := 0

	for row < len(input) {
		if input[row][col%len(input[row])] == '#' {
			count++
		}

		row += 1
		col += 3
	}

	return count
}
