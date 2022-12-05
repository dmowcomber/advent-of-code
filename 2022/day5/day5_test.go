package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const day = 5

func TestPart1Example(t *testing.T) {
	count, err := getPart1("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day%d example answer: %s", day, count)
	assert.Equal(t, "CMZ", count)
}

func TestPart1(t *testing.T) {
	t.Skip()
	count, err := getPart1("./input.txt")
	require.NoError(t, err)

	t.Logf("day%d answer: %s", day, count)
	assert.Equal(t, "GFTNRBZPF", count)
}

func TestPart2Example(t *testing.T) {
	count, err := getPart2("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day%d example answer: %s", day, count)
	assert.Equal(t, "MCD", count)
}

func TestPart2(t *testing.T) {
	count, err := getPart2("./input.txt")
	require.NoError(t, err)

	t.Logf("day%d answer: %s", day, count)
	assert.Equal(t, "VRQWPDSGP", count)
}
