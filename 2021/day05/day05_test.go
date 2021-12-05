package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{}

func TestPart1(t *testing.T) {
	result := SolvePart1(exampleInput)

	assert.Equal(t, 99999, result)
}

func TestPart2(t *testing.T) {
	result := SolvePart2(exampleInput)

	assert.Equal(t, 99999, result)
}
