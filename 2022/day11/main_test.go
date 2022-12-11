package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	answer, err := getMonkeyBusinessLevel("./input-example.txt", 20, true)
	require.NoError(t, err)

	t.Logf("part1 example answer: %d", answer)
	assert.Equal(t, 10605, answer)
}

func TestPart1(t *testing.T) {
	answer, err := getMonkeyBusinessLevel("./input.txt", 20, true)
	require.NoError(t, err)

	t.Logf("part1 answer: %d", answer)
	assert.Equal(t, 120056, answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := getMonkeyBusinessLevel("./input-example.txt", 10000, false)
	require.NoError(t, err)

	t.Logf("part2 example answer: %d", answer)
	assert.Equal(t, 2713310158, answer)
}

func TestPart2(t *testing.T) {
	answer, err := getMonkeyBusinessLevel("./input.txt", 10000, false)
	require.NoError(t, err)

	t.Logf("part2 answer: %d", answer)
	assert.Equal(t, 21816744824, answer)
}
