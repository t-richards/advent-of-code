package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestPart1(t *testing.T) {
	actual := SolvePart1(exampleInput)

	assert.Equal(t, 150, actual)
}

func TestPart2(t *testing.T) {
	actual := SolvePart2(exampleInput)

	assert.Equal(t, 900, actual)
}
