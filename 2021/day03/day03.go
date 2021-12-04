package main

import (
	"aoc/internal/conv"
	"aoc/internal/data"
	"fmt"
	"strings"
)

func main() {
	input := data.LoadStrings()

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

func SolvePart1(input []string) int {
	gamma := 0
	epsilon := 0

	threshold := len(input) / 2
	transposed := conv.Transpose(input)
	transposedLen := len(transposed)

	for i, line := range transposed {
		if strings.Count(line, "0") >= threshold {
			gamma |= 1 << (transposedLen - 1 - i)
		} else {
			epsilon |= 1 << (transposedLen - 1 - i)
		}
	}

	return gamma * epsilon
}

func SolvePart2(input []string) int64 {
	return 0
}
