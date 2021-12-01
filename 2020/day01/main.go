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

	fmt.Println(ExpensePart1(input))
	fmt.Println(ExpensePart2(input))
}

func ExpensePart1(entries []int) int {
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i]+entries[j] == 2020 {
				return entries[i] * entries[j]
			}
		}
	}

	return -1
}

func ExpensePart2(entries []int) int {
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			for k := i + 2; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == 2020 {
					return entries[i] * entries[j] * entries[k]
				}
			}
		}
	}

	return -1
}
