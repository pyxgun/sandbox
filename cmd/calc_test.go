package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actulal := sum(3, 5)
	expected := 8

	assert.Equal(t, expected, actulal)
}

func TestSub(t *testing.T) {
	actual := sub(5, 1)
	expected := 3

	assert.Equal(t, expected, actual)
}
