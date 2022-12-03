package main

import (
	"errors"
	"fmt"
)

const (
	strategyLose = "X"
	strategyDraw = "Y"
	strategyWin  = "Z"
)

func getScorePart2(filename string) (int, error) {
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
		strategy := string(line[2])
		var score int

		if strategy == strategyDraw {
			score = draw
			switch elf {
			case rockElf:
				score += rockScore
			case paperElf:
				score += paperScore
			case scissorsElf:
				score += scissorsScore
			default:
				return 0, errors.New("unexpected elf choice")
			}
		} else if strategy == strategyWin {
			score = win
			switch elf {
			case rockElf:
				score += paperScore
			case paperElf:
				score += scissorsScore
			case scissorsElf:
				score += rockScore
			default:
				return 0, errors.New("unexpected elf choice")
			}
		} else if strategy == strategyLose {
			score = lose
			switch elf {
			case rockElf:
				score += scissorsScore
			case paperElf:
				score += rockScore
			case scissorsElf:
				score += paperScore
			default:
				return 0, errors.New("unexpected elf choice")
			}
		}
		totalScore += score
	}
	return totalScore, nil
}
