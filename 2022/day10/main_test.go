package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Example2(t *testing.T) {
	answer, err := getPart1("./input-example.txt")
	require.NoError(t, err)

	t.Logf("part1 example answer: %d", answer)
	assert.Equal(t, 13140, answer)
}

func TestPart1(t *testing.T) {
	answer, err := getPart1("./input.txt")
	require.NoError(t, err)

	t.Logf("part1 answer: %d", answer)
	assert.Equal(t, 14860, answer)
}

func TestPart2Example(t *testing.T) {
	answer, err := getPart2("./input-example.txt")
	require.NoError(t, err)

	t.Logf("part2 example answer: %s", answer)

	expected := `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	assert.Equal(t, expected, answer)
}

func TestPart2(t *testing.T) {
	answer, err := getPart2("./input.txt")
	require.NoError(t, err)

	t.Logf("part2 answer: %s", answer)

	// RGZEHURK
	expected := `
###...##..####.####.#..#.#..#.###..#..#.
#..#.#..#....#.#....#..#.#..#.#..#.#.#..
#..#.#......#..###..####.#..#.#..#.##...
###..#.##..#...#....#..#.#..#.###..#.#..
#.#..#..#.#....#....#..#.#..#.#.#..#.#..
#..#..###.####.####.#..#..##..#..#.#..#.`
	assert.Equal(t, expected, answer)
}
