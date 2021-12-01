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
	fmt.Println(ValidPasswords2(input))
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
			result++
		}
	}

	return result
}

func ValidPasswords2(lines []string) int {
	result := 0

	for _, line := range lines {
		found := 0
		i, j, letter, password := parsePolicy(line)
		if password[i-1] == letter[0] {
			found += 1
		}
		if password[j-1] == letter[0] {
			found += 1
		}

		if found == 1 {
			result++
		}
	}

	return result
}
