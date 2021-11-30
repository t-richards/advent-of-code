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

	fmt.Println(Expense(input))
}

type CombinationIterator struct {
	n              int
	size           int
	firstIteration bool
	pool           []int
	indices        []int
}

func Expense(entries []int) int {
	iterator := NewCombinationIterator(entries, 2)

	for iterator.Next() {

		if iterator.Sum() != 2020 {
			continue
		}

		return iterator.Product()
	}

	return -1
}

func NewCombinationIterator(s []int, size int) *CombinationIterator {
	iter := &CombinationIterator{
		n:              len(s),
		size:           size,
		firstIteration: true,
		indices:        make([]int, len(s)),
		pool:           make([]int, len(s)),
	}

	if size > iter.n {
		return nil
	}

	for i := 0; i < len(s); i++ {
		iter.indices[i] = i
	}

	copy(iter.pool, s)

	return iter
}

func (c *CombinationIterator) Next() bool {
	if c.firstIteration {
		c.firstIteration = false
		return true
	}

	stop := true
	i := c.size - 1
	for i >= 0 {
		if c.indices[i] != i+c.n-c.size {
			stop = false
			break
		}
		i -= 1
	}

	if stop {
		return false
	}

	c.indices[i] += 1
	c.pool[i] = c.pool[c.indices[i]]

	for j := i + 1; j < c.size-1; j++ {
		c.indices[j] = c.indices[j-1] + 1
		c.pool[j] = c.pool[c.indices[j]]
	}

	return true
}

func (c CombinationIterator) Sum() int {
	sum := 0
	for _, v := range c.pool[:c.size] {
		sum += v
	}
	return sum
}

func (c CombinationIterator) Product() int {
	sum := 1
	for _, v := range c.pool[:c.size] {
		sum *= v
	}
	return sum
}
