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

func toi2(input string) int64 {
	value, _ := strconv.ParseInt(input, 2, 64)

	return value
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

// Slightly janky string stuff here
func SolvePart1(input []string) int64 {
	gamma := ""
	epsilon := ""

	threshold := len(input) / 2
	transposed := transpose(input)

	for _, line := range transposed {
		num := toi2(line)
		if bits.OnesCount64(uint64(num)) >= threshold {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	return toi2(gamma) * toi2(epsilon)
}

func SolvePart2(input []string) int64 {
	return 0
}
