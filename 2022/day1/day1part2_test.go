package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1Part2(t *testing.T) {
	sumTop3Calories, err := getSumTop3Caloies()
	require.NoError(t, err)

	t.Logf("day2 answer: %d", sumTop3Calories)
}
