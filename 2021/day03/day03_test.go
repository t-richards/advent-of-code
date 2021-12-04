package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1(t *testing.T) {
	actual := SolvePart1(exampleInput)

	assert.Equal(t, int64(198), actual)
}

func TestPart2(t *testing.T) {
	actual := SolvePart2(exampleInput)

	assert.Equal(t, 230, actual)
}
