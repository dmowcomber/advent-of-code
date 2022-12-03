package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// elfs choice
	rockElf     = "A"
	paperElf    = "B"
	scissorsElf = "C"

	// your choice
	rockMe     = "X"
	paperMe    = "Y"
	scissorsMe = "Z"

	// scored by your choice
	rockScore     = 1
	paperScore    = 2
	scissorsScore = 3

	win  = 6
	draw = 3
	lose = 0
)

func getScore(filename string) (int, error) {
	lines, err := readLines(filename)
	if err != nil {
		return 0, err
	}

	var totalScore int
	for _, line := range lines {
		if len(line) < 3 {
			return 0, fmt.Errorf("unexpected input size: %q. should be at least 3 characters", line)
		}
		elf := string(line[0])
		me := string(line[2])
		winLoseTieScore := getWinLoseTieScore(elf, me)
		choiceScore, err := getChoiceScore(me)
		if err != nil {
			return 0, err
		}
		score := winLoseTieScore + choiceScore
		totalScore += score
	}

	return totalScore, nil
}

func getChoiceScore(me string) (int, error) {
	if me == rockMe {
		return rockScore, nil
	}
	if me == paperMe {
		return paperScore, nil
	}
	if me == scissorsMe {
		return scissorsScore, nil
	}
	return 0, fmt.Errorf("unexepcted rock/paper/scissors choice: %q", me)
}

func getWinLoseTieScore(elf, me string) int {
	// Rock beats scissors
	// Paper beats rock
	// Scissors beats paper

	if me == rockMe && elf == scissorsElf ||
		me == paperMe && elf == rockElf ||
		me == scissorsMe && elf == paperElf {
		return win
	}

	if elf == rockElf && me == scissorsMe ||
		elf == paperElf && me == rockMe ||
		elf == scissorsElf && me == paperMe {
		return lose
	}

	return draw
}

func readLines(filename string) ([]string, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	result := make([]string, 0, 100)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result, nil
}
