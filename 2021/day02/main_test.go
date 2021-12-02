package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []int{1, 2, 3}

func TestPart1(t *testing.T) {
	actual := SolvePart1(exampleInput)

	assert.Equal(t, 7, actual)
}

func TestPart2(t *testing.T) {
	actual := SolvePart2(exampleInput)

	assert.Equal(t, 5, actual)
}
