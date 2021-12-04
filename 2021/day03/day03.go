package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
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

func transpose(input []string) []string {
	columnLen := len(input[0])
	result := make([]string, columnLen)

	for col := 0; col < columnLen; col++ {
		for row := 0; row < len(input); row++ {
			result[col] += string(input[row][col])
		}
	}

	return result
}

func SolvePart1(input []string) int {
	gamma := 0
	epsilon := 0

	threshold := len(input) / 2
	transposed := transpose(input)

	for i, line := range transposed {
		num, _ := strconv.ParseUint(line, 2, 64)
		if bits.OnesCount64(num) >= threshold {
			gamma |= 1 << (len(transposed) - 1 - i)
		} else {
			epsilon |= 1 << (len(transposed) - 1 - i)
		}
	}

	return gamma * epsilon
}

func SolvePart2(input []string) int64 {
	return 0
}
