package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	answer, err := getLastTailPointCount("./input-example.txt", 1)
	require.NoError(t, err)

	t.Logf("part1 example answer: %d", answer)
	assert.Equal(t, 13, answer)
}

func TestPart1(t *testing.T) {
	answer, err := getLastTailPointCount("./input.txt", 1)
	require.NoError(t, err)

	t.Logf("part1 answer: %d", answer)
	assert.Equal(t, 6271, answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := getLastTailPointCount("./input-example.txt", 9)
	require.NoError(t, err)

	t.Logf("part2 example answer: %d", answer)
	assert.Equal(t, 1, answer)
}

func TestPart2Example2(t *testing.T) {
	answer, err := getLastTailPointCount("./input-example2.txt", 9)
	require.NoError(t, err)

	t.Logf("part2 example answer: %d", answer)
	assert.Equal(t, 36, answer)
}

func TestPart2(t *testing.T) {
	answer, err := getLastTailPointCount("./input.txt", 9)
	require.NoError(t, err)

	t.Logf("part2 answer: %d", answer)
	assert.Equal(t, 2458, answer)
}
