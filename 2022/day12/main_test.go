package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	answer, err := getCount("./input-example.txt", 'S')
	require.NoError(t, err)

	t.Logf("part1 example answer: %d", answer)
	assert.Equal(t, 31, answer)
}

func TestPart1(t *testing.T) {
	answer, err := getCount("./input.txt", 'S')
	require.NoError(t, err)

	t.Logf("part1 answer: %d", answer)
	assert.Equal(t, 423, answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := getCount("./input-example.txt", 'a')
	require.NoError(t, err)

	t.Logf("part2 example answer: %d", answer)
	assert.Equal(t, 29, answer)
}

func TestPart2(t *testing.T) {
	answer, err := getCount("./input.txt", 'a')
	require.NoError(t, err)

	t.Logf("part2 answer: %d", answer)
	assert.Equal(t, 416, answer)
}
