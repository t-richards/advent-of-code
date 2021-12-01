package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

func SolvePart1(input []int) int {
	count := 0

	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			count++
		}
	}

	return count
}

func SolvePart2(input []int) int {
	count := 0

	for i := 3; i < len(input); i++ {
		prevWindow := input[i-1] + input[i-2] + input[i-3]
		currentWindow := input[i] + input[i-1] + input[i-2]

		if prevWindow < currentWindow {
			count++
		}
	}

	return count
}
