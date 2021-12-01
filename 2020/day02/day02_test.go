package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPart1(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := ValidPasswords1(input)

	assert.Equal(t, 2, result)
}
