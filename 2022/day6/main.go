package main

import "github.com/dmowcomber/advent-of-code/input"

func uniqueIndex(nLetters int, filename string) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	lastXIndex := 0
	lastXLetters := make([]rune, nLetters)
	for i, letter := range lines[0] { // there's always one line in this input
		// keep the last X numbers only
		lastXLetters[lastXIndex] = letter
		lastXIndex = (lastXIndex + 1) % nLetters

		// check unique count of last X letters
		uniqueLetters := make(map[rune]struct{})
		for index := 0; index < len(lastXLetters) && lastXLetters[index] != 0; index++ {
			uniqueLetters[lastXLetters[index]] = struct{}{}
		}
		if len(uniqueLetters) == nLetters {
			return i + 1, nil // if all letters are unique, return
		}
	}
	return -1, nil
}
