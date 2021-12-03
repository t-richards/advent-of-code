package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Janky string version, not sorry
func SolvePart1(input []string) int {
	var gamma strings.Builder
	var epsilon strings.Builder

	for col := 0; col < len(input[0]); col++ {
		zeroCnt := 0
		oneCnt := 0

		for row := 0; row < len(input); row++ {
			if input[row][col] == '0' {
				zeroCnt++
			} else {
				oneCnt++
			}
		}

		if zeroCnt > oneCnt {
			gamma.WriteRune('1')
			epsilon.WriteRune('0')
		} else {
			gamma.WriteRune('0')
			epsilon.WriteRune('1')
		}
	}

	gam, _ := strconv.ParseInt(gamma.String(), 2, 64)
	eps, _ := strconv.ParseInt(epsilon.String(), 2, 64)

	return int(gam * eps)
}

func SolvePart2(input []string) int64 {
	return 0
}
