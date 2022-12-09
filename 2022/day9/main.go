package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dmowcomber/advent-of-code/input"
)

func getLastTailPointCount(filename string, maxPointIndex int) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	// get motions input
	motions := make([]motion, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return 0, fmt.Errorf("unexpected length of input: %d", len(parts))
		}

		direction := parts[0]
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, fmt.Errorf("unexpected step input: %s", parts[1])
		}

		motions = append(motions, motion{
			direction: direction,
			steps:     steps,
		})
	}

	type point struct {
		x, y int
	}
	type connectedPoint struct {
		point point
		next  *connectedPoint
	}
	head := &connectedPoint{}

	// track visited points
	visited := make(map[point]struct{})

	for _, mot := range motions {
		// log.Printf("motion: %#v", mot)
		for i := 0; i < mot.steps; i++ {
			// update head position
			switch mot.direction {
			case "U":
				head.point.y++
			case "D":
				head.point.y--
			case "L":
				head.point.x--
			case "R":
				head.point.x++
			}

			// check tail position vs head position and update tail
			current := head
			for pointIndex := 0; pointIndex < maxPointIndex; pointIndex++ {
				if current.next == nil {
					current.next = &connectedPoint{}
				}
				tail := current.next

				dy := abs(current.point.y - tail.point.y)
				dx := abs(current.point.x - tail.point.x)

				// pull tail closer if tail is too far up
				if current.point.y-tail.point.y <= -2 {
					tail.point.y = current.point.y + 1
					if dx == 1 {
						tail.point.x = current.point.x
					}
				}
				// pull tail closer if tail is too far down
				if current.point.y-tail.point.y >= 2 {
					tail.point.y = current.point.y - 1
					if dx == 1 {
						tail.point.x = current.point.x
					}
				}
				// pull tail closer if tail is too far left
				if current.point.x-tail.point.x >= 2 {
					tail.point.x = current.point.x - 1
					// tail.point.x++
					if dy == 1 {
						tail.point.y = current.point.y
					}
				}
				// pull tail closer if tail is too far right
				if current.point.x-tail.point.x <= -2 {
					tail.point.x = current.point.x + 1
					if dy == 1 {
						tail.point.y = current.point.y
					}
				}

				if pointIndex+1 == maxPointIndex {
					visited[tail.point] = struct{}{}
					break
				}
				current = tail
			}
		}
	}
	return len(visited), nil
}

type motion struct {
	direction string
	steps     int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
