package main

import (
	"strconv"
	"strings"

	"github.com/dmowcomber/advent-of-code/input"
)

func getCount(filename string) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, line := range lines {
		a, b, x, y, err := getNums(line)
		if err != nil {
			return 0, nil
		}
		// log.Printf("a,b,x,y: %v %v %v %v", a, b, x, y)

		// 6-6,4-7
		// a-b,x-y
		if a <= x && y <= b {
			count++
		} else if x <= a && b <= y {
			count++
		}
	}

	return count, nil
}

func getNums(line string) (int, int, int, int, error) {

	parts := strings.Split(line, ",")
	parts1 := strings.Split(parts[0], "-")
	parts2 := strings.Split(parts[1], "-")

	a, err := strconv.Atoi(parts1[0])
	if err != nil {
		return 0, 0, 0, 0, nil
	}
	b, err := strconv.Atoi(parts1[1])
	if err != nil {
		return 0, 0, 0, 0, nil
	}
	x, err := strconv.Atoi(parts2[0])
	if err != nil {
		return 0, 0, 0, 0, nil
	}
	y, err := strconv.Atoi(parts2[1])
	if err != nil {
		return 0, 0, 0, 0, nil
	}
	return a, b, x, y, nil
}
