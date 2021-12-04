package main

import (
	"aoc/internal/conv"
	"aoc/internal/data"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := data.LoadStrings()

	fmt.Println(SolvePart1(input))
	fmt.Println(SolvePart2(input))
}

func PowerRates(input []string) (string, string) {
	var gamma strings.Builder
	var epsilon strings.Builder

	threshold := len(input) / 2
	transposed := conv.Transpose(input)

	for _, line := range transposed {
		if strings.Count(line, "1") >= threshold {
			gamma.WriteRune('1')
			epsilon.WriteRune('0')
		} else {
			epsilon.WriteRune('1')
			gamma.WriteRune('0')
		}
	}

	return gamma.String(), epsilon.String()
}

func LifeRates(input []string) ([]int, []int) {
	zeroesCount := make([]int, len(input[0]))
	onesCount := make([]int, len(input[0]))

	for _, line := range input {
		for i, char := range line {
			if char == '1' {
				onesCount[i]++
			} else {
				zeroesCount[i]++
			}
		}
	}

	return zeroesCount, onesCount
}

func keep(s []string, idx int, r rune) []string {
	if len(s) == 1 {
		return s
	}

	newlist := []string{}

	for _, item := range s {
		if rune(item[idx]) == r {
			newlist = append(newlist, item)
		}
	}

	return newlist
}

func SolvePart1(input []string) int {
	a, b := PowerRates(input)
	gamma, _ := strconv.ParseInt(a, 2, 64)
	epsilon, _ := strconv.ParseInt(b, 2, 64)

	return int(gamma * epsilon)
}

func SolvePart2(input []string) int {
	columns := len(input[0])
	oxylist := make([]string, len(input))
	copy(oxylist, input)
	co2list := make([]string, len(input))
	copy(co2list, input)

	for i := 0; i < columns; i++ {
		zeroes, ones := LifeRates(oxylist)
		if ones[i] >= zeroes[i] {
			oxylist = keep(oxylist, i, '1')
		} else {
			oxylist = keep(oxylist, i, '0')
		}
	}

	for i := 0; i < columns; i++ {
		zeroes, ones := LifeRates(co2list)
		if zeroes[i] <= ones[i] {
			co2list = keep(co2list, i, '0')
		} else {
			co2list = keep(co2list, i, '1')
		}
	}

	oxygen, _ := strconv.ParseInt(oxylist[0], 2, 64)
	co2, _ := strconv.ParseInt(co2list[0], 2, 64)

	return int(oxygen * co2)
}
