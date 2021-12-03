package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{}

func TestPart1(t *testing.T) {
	actual := SolvePart1(exampleInput)

	assert.Equal(t, 1111111, actual)
}

func TestPart2(t *testing.T) {
	actual := SolvePart2(exampleInput)

	assert.Equal(t, 22222222, actual)
}
