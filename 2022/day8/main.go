package main

import (
	"strconv"

	"github.com/dmowcomber/advent-of-code/input"
)

func getPart1(filename string) (int, error) {
	trees, err := getInput(filename)
	if err != nil {
		return 0, err
	}

	var length, outerLength int
	for i, treeRowCol := range trees {
		length = len(treeRowCol)
		outerLength = len(trees)

		blockedLine := -1
		for j, tree := range treeRowCol {
			blockedLine = checkVisibility(i, j, outerLength, length, tree, blockedLine)
		}
		blockedLine = -1
		for j := len(treeRowCol) - 1; j >= 0; j-- {
			tree := treeRowCol[j]
			blockedLine = checkVisibility(i, j, outerLength, length, tree, blockedLine)
		}
	}

	for i := 0; i < length; i++ {
		blockedLine := -1
		for j := 0; j < outerLength; j++ {
			tree := trees[j][i]
			blockedLine = checkVisibility(i, j, outerLength, length, tree, blockedLine)

		}
		blockedLine = -1
		for j := outerLength - 1; j >= 0; j-- {
			tree := trees[j][i]
			blockedLine = checkVisibility(i, j, outerLength, length, tree, blockedLine)

		}
	}

	// get count
	count := 0
	for _, treeRowCol := range trees {
		for _, tree := range treeRowCol {
			if tree.isVisible {
				count++
			}
		}
	}

	return count, nil
}

func checkVisibility(outerIndex, index, outerLength, length int, tree *treeVisibility, blockedLine int) int {
	if index == 0 || outerIndex == 0 || index == length-1 || outerIndex == outerLength-1 {
		tree.isVisible = true
		return tree.height
	}

	if tree.height > blockedLine {
		tree.isVisible = true
		return tree.height
	}
	return blockedLine
}

type treeVisibility struct {
	height    int
	isVisible bool
	score     int
}

func getPart2(filename string) (int, error) {
	trees, err := getInput(filename)
	if err != nil {
		return 0, err
	}

	// get scores
	var length, outerLength int
	for i, treeRowCol := range trees {
		length = len(treeRowCol)
		outerLength = len(trees)

		for j, tree := range treeRowCol {
			checkScore(i, j, outerLength, length, tree, trees)
		}
	}

	// return highest score
	highScore := 0
	for _, treeRowCol := range trees {
		for _, tree := range treeRowCol {
			if tree.score > highScore {
				highScore = tree.score
			}
		}
	}

	return highScore, nil
}

func checkScore(outerIndex, index, outerLength, length int, tree *treeVisibility, trees [][]*treeVisibility) {
	if outerIndex == 0 || outerIndex == outerLength-1 {
		tree.score = 0
		return
	}
	if index == 0 || index == length-1 {
		tree.score = 0
		return
	}

	scoreUp := 0
	for outerI := outerIndex - 1; outerI >= 0; outerI-- {
		scoreUp++
		if tree.height <= trees[outerI][index].height {
			break
		}
	}
	score := scoreUp

	scoreDown := 0
	for outerI := outerIndex + 1; outerI < outerLength; outerI++ {
		scoreDown++
		if tree.height <= trees[outerI][index].height {
			break
		}
	}
	score = score * scoreDown

	scoreLeft := 0
	for innderI := index - 1; innderI >= 0; innderI-- {
		scoreLeft++
		if tree.height <= trees[outerIndex][innderI].height {
			break
		}
	}
	score = score * scoreLeft

	scoreRight := 0
	for innderI := index + 1; innderI < length; innderI++ {
		scoreRight++
		if tree.height <= trees[outerIndex][innderI].height {
			break
		}
	}
	score = score * scoreRight

	tree.score = score
}

func getInput(filename string) ([][]*treeVisibility, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return nil, err
	}

	trees := make([][]*treeVisibility, 0)
	for i, line := range lines {
		treesRow := make([]*treeVisibility, 0)
		trees = append(trees, treesRow)

		for _, s := range line {
			treeHeight, err := strconv.Atoi(string(s))
			if err != nil {
				return nil, err
			}

			treeVis := &treeVisibility{
				height: treeHeight,
			}

			treesRow = append(treesRow, treeVis)
		}
		trees[i] = treesRow
	}
	return trees, nil
}
