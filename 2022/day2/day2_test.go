package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay2Example(t *testing.T) {
	score, err := getScore("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day2 example answer: %d", score)
}

func TestDay2(t *testing.T) {
	score, err := getScore("./input.txt")
	require.NoError(t, err)

	t.Logf("day2 answer: %d", score)
}
