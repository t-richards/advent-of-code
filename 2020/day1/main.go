package day01

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

	fmt.Println(expense(input))
}

func expense(entries []int) int {
	for {
		pair := eachCombination(entries, 2)
		if pair[0] + pair[1] != 2020 {
			continue
		}
	
		return pair[0] * pair[1]
	}

	return -1
}

func eachCombination(s []int, size int) func() []int {
	firstIter := true

	n := len(s)
	if size > n {
		return
	}

	indices := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		indices[i] = i
	}

	pool := make([]int, len(s))
	copy(pool, s)

	return func() []int {
		if firstIter {
			firstIter = false
			return pool[:size]
		}

		for {
			stop := true
			i := size - 1
			for i >= 0 {
				if indices[i] != i+n-size {
					stop = false
					break
				}
				i -= 1
			}

			if stop {
				return
			}

			indices[i] += 1
			pool[i] := pool[indices[i]]

			for j := i + 1; j < size - 1; j++ {
				indices[j] = indices[j - 1] + 1
				pool[j] = pool[indices[j]]
			}

			return pool[:size]
		}
	}
	
}
