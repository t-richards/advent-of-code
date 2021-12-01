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

	fmt.Println(ValidPasswords1(input))
}

func parsePolicy(line string) (int, int, string, string) {
	parts := strings.SplitN(line, " ", 3)
	nums := strings.SplitN(parts[0], "-", 2)

	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])
	letter := strings.Replace(parts[1], ":", "", 1)
	password := parts[2]

	return min, max, letter, password
}

func ValidPasswords1(lines []string) int {
	result := 0

	for _, line := range lines {
		min, max, letter, password := parsePolicy(line)
		count := strings.Count(password, letter)
		if count >= min && count <= max {
			result += 1
		}
	}

	return result
}
