package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetExample(t *testing.T) {
	priority, err := getPrirority("./input-example.txt")
	require.NoError(t, err)

	t.Logf("day3 example answer: %d", priority)
}

func TestGetPriority(t *testing.T) {
	priority, err := getPrirority("./input.txt")
	require.NoError(t, err)

	t.Logf("day3 answer: %d", priority)
}
