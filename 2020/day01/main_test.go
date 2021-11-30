package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpense(t *testing.T) {
	input := []int{1721, 979, 366, 299, 675, 1456}

	actual := Expense(input)

	assert.Equal(t, actual, 514579)
}
