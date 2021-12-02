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
	return 0
}

func SolvePart2(input []int) int {
	return 0
}
