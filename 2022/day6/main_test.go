package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	answer, err := uniqueIndex(4, "./input-example.txt")
	require.NoError(t, err)

	t.Logf("part1 example answer: %d", answer)
	assert.Equal(t, 7, answer)
}

func TestPart1(t *testing.T) {
	answer, err := uniqueIndex(4, "./input.txt")
	require.NoError(t, err)

	t.Logf("part1 answer: %d", answer)
	assert.Equal(t, 1175, answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := uniqueIndex(14, "./input-example.txt")
	require.NoError(t, err)

	t.Logf("part2 example answer: %d", answer)
	assert.Equal(t, 19, answer)
}

func TestPart2(t *testing.T) {
	answer, err := uniqueIndex(14, "./input.txt")
	require.NoError(t, err)

	t.Logf("part2 answer: %d", answer)
	assert.Equal(t, 3217, answer)
}
