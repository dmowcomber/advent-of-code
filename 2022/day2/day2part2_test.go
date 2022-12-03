package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay2Part2Example(t *testing.T) {
	score, err := getScorePart2("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day2 part2 example answer: %d", score)
}

func TestDay2Part2(t *testing.T) {
	score, err := getScorePart2("./input.txt")
	require.NoError(t, err)

	t.Logf("day2 part2 answer: %d", score)
}
