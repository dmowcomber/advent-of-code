package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	count, err := getCount("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day4 example answer: %d", count)
	assert.Equal(t, 2, count)
}

func TestPart1(t *testing.T) {
	count, err := getCount("./input.txt")
	require.NoError(t, err)

	t.Logf("day4 answer: %d", count)
	assert.Equal(t, 448, count)
}

func TestPart2Example(t *testing.T) {
	count, err := getCountPart2("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day4 example answer: %d", count)
	assert.Equal(t, 4, count)
}

func TestPart2(t *testing.T) {
	count, err := getCountPart2("./input.txt")
	require.NoError(t, err)

	t.Logf("day4 answer: %d", count)
	assert.Equal(t, 794, count)
}
