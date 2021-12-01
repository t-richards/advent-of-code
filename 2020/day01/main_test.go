package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	actual := ExpensePart1(input)

	assert.Equal(t, actual, 514579)
}

func TestPart2(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	actual := ExpensePart2(input)

	assert.Equal(t, actual, 241861950)
}
