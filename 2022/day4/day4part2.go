package main

import (
	"github.com/dmowcomber/advent-of-code/input"
)

func getCountPart2(filename string) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	var count int
	for _, line := range lines {
		a, b, x, y, err := getNums(line)
		if err != nil {
			return 0, nil
		}
		// log.Printf("a,b,x,y: %v %v %v %v", a, b, x, y)

		// examples:
		// 1-5,4-7
		// a-b,x-y
		// 1<=4,5>=4
		// 5<=6,6<=6

		// 6-6,4-6
		// a<=y,b>=y

		// 4-6,6-6
		// a<=6,6>=6

		if inRangeInclusive(a, x, y) ||
			inRangeInclusive(b, x, y) ||
			inRangeInclusive(x, a, b) ||
			inRangeInclusive(y, a, b) {
			count++
		}
	}

	return count, nil
}

// inRangeInclusive returns true is i is in the range of x-y inclusive
func inRangeInclusive(i, x, y int) bool {
	return x <= i && i <= y
}
