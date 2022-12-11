package main

import (
	"strconv"
	"strings"

	"github.com/dmowcomber/advent-of-code/input"
)

func getPart1(filename string) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	x := 1
	cycle := 0
	mod := 20

	sigStrengths := 0

	for _, line := range lines {
		if line == "noop" {
			cycle++

			if cycle == mod {
				sigStrengths += mod * x
				mod += 40
			}

			continue
		}

		parts := strings.Split(line, " ")
		numStr := parts[1]
		num, _ := strconv.Atoi(numStr)
		// log.Printf("num: %d", num)

		cycle++

		if cycle == mod {
			sigStrengths += mod * x
			mod += 40
		}
		cycle++

		if cycle == mod {
			sigStrengths += mod * x
			mod += 40

		}

		x += num
	}

	// log.Printf("last cycle: %d", cycle)
	return sigStrengths, nil
}

func getPart2(filename string) (string, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return "", err
	}

	x := 1
	cycle := 0

	xList := make([]int, 0)
	for _, line := range lines {
		if line == "noop" {
			cycle++
			xList = append(xList, x)
			continue
		}

		parts := strings.Split(line, " ")
		numStr := parts[1]
		num, _ := strconv.Atoi(numStr)

		cycle++
		xList = append(xList, x)

		cycle++
		xList = append(xList, x)
		x += num
	}

	// lazy string building
	result := ""
	for i, x := range xList {
		position := i % 40
		if position == 0 {
			result += "\n"
		}
		if x == position-1 || position == x || x == position+1 {
			result += "#"
		} else {
			result += "."
		}
	}
	return result, nil
}
