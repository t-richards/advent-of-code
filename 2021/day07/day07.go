package main

import (
	"aoc/internal/data"
	"fmt"
	"math"
)

func minmax(s []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	for _, n := range s {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min, max
}

func abs(i int) int {
	v := math.Abs(float64(i))
	return int(v)
}

func main() {
	input := data.LoadSplitInts()

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

func SolvePart1(input []int) int {
	lowest := math.MaxInt
	min, max := minmax(input)

	for i := min; i <= max; i++ {
		fuel := 0

		for _, crab := range input {
			fuel += abs(crab - i)
		}

		if fuel < lowest {
			lowest = fuel
		}
	}

	return lowest
}

func SolvePart2(input []int) int {
	lowest := math.MaxInt
	min, max := minmax(input)

	for i := min; i <= max; i++ {
		fuel := 0

		for _, crab := range input {
			n := abs(crab - i)
			fuel += (n * (n + 1)) / 2
		}

		if fuel < lowest {
			lowest = fuel
		}
	}

	return lowest
}
