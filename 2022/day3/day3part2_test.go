package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPriorityPart2Example(t *testing.T) {
	priority, err := getPrirorityPart2("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day3 part2 example answer: %d", priority)
}

func TestGetPriorityPart2(t *testing.T) {
	// t.Skip()
	priority, err := getPrirorityPart2("./input.txt")
	require.NoError(t, err)

	t.Logf("day3 part2 answer: %d", priority)
}
