package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []int{1, 2, 3}

	actual := SolvePart1(input)

	assert.Equal(t, -1, actual)
}

func TestPart2(t *testing.T) {
	input := []int{4, 5, 6}

	actual := SolvePart2(input)

	assert.Equal(t, -1, actual)
}
