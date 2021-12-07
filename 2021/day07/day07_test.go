package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []int{
	16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
}

func TestPart1(t *testing.T) {
	result := SolvePart1(exampleInput)

	assert.Equal(t, 37, result)
}

func TestPart2(t *testing.T) {
	result := SolvePart2(exampleInput)

	assert.Equal(t, 168, result)
}
