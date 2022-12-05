package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example(t *testing.T) {
	answer, err := getPart1("./input-example.txt")
	require.NoError(t, err)

	t.Logf("example answer: %s", answer)
	assert.Equal(t, "CMZ", answer)
}

func TestPart1(t *testing.T) {
	answer, err := getPart1("./input.txt")
	require.NoError(t, err)

	t.Logf("answer: %s", answer)
	assert.Equal(t, "GFTNRBZPF", answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := getPart2("./input-example.txt")
	require.NoError(t, err)

	t.Logf("part2 example answer: %s", answer)
	assert.Equal(t, "MCD", answer)
}

func TestPart2(t *testing.T) {
	answer, err := getPart2("./input.txt")
	require.NoError(t, err)

	t.Logf("part2 answer: %s", answer)
	assert.Equal(t, "VRQWPDSGP", answer)
}
