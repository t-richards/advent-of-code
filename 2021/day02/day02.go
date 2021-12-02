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

func SolvePart1(input []string) int {
	hor := 0
	dep := 0

	for _, line := range input {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			hor += amount
		case "down":
			dep += amount
		case "up":
			dep -= amount
		}
	}

	return hor * dep
}

func SolvePart2(input []string) int {
	hor := 0
	dep := 0
	aim := 0

	for _, line := range input {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "down":
			aim += amount
		case "up":
			aim -= amount
		case "forward":
			hor += amount
			dep += aim * amount
		}
	}

	return hor * dep
}
