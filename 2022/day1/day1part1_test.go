package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	maxCalories, err := getMaxCaloies()
	require.NoError(t, err)

	t.Logf("maxCalories: %d", maxCalories)
}
