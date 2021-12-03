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

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	aaaaaaa := 0

	return aaaaaaa
}

func SolvePart2(input []string) int {
	bbbbbbbb := 0

	return bbbbbbbb
}
