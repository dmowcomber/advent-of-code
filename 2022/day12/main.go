package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/dmowcomber/advent-of-code/input"
)

type point struct{ y, x int }

type traverser struct {
	lines         []string
	visited       map[point]bool
	nextPositions map[point]struct{}
	positions     map[point]struct{}
}

func getCount(filename string, startLetter rune) (int, error) {
	debug := false
	logger := debugLogger{
		debug: debug,
	}
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	terrain := &traverser{
		lines:         lines,
		visited:       make(map[point]bool),
		positions:     make(map[point]struct{}),
		nextPositions: make(map[point]struct{}),
	}

	count := 0
	for j, line := range lines {
		for i, r := range line {
			if r == startLetter {
				startPosition := point{y: j, x: i}
				terrain.positions[startPosition] = struct{}{}
				terrain.positions[startPosition] = struct{}{}
			}
		}
	}

	for ; ; count++ {
		logger.Printf("count: %d", count)
		if debug {
			terrain.print()
		}

		if len(terrain.positions) == 0 {
			return 0, errors.New("unexpected positions count")
		}

		// attempt all positions all at once until we reach E
		for pos := range terrain.positions {
			x, y := pos.x, pos.y
			letter := lines[y][x]

			// logger.Printf("positions: %#v", positions)
			// logger.Printf("pos: %v, letter: %s", pos, string(letter))

			if letter == 'E' {
				return count, nil
			}

			// up
			if y != 0 {
				terrain.attemptTraverse(y-1, x, lines, letter)
			}
			// down
			if y < len(lines)-1 {
				terrain.attemptTraverse(y+1, x, lines, letter)
			}
			// left
			if x != 0 {
				terrain.attemptTraverse(y, x-1, lines, letter)
			}
			// right
			if x < len(lines[y])-1 {
				terrain.attemptTraverse(y, x+1, lines, letter)
			}
		}
		terrain.positions = terrain.nextPositions
		terrain.nextPositions = make(map[point]struct{})
	}
}

func (t *traverser) attemptTraverse(nextY, nextX int, lines []string, letter byte) {
	letterNext := lines[nextY][nextX]
	if letter == 'S' || rune(letterNext)-rune(letter) <= 1 {
		// don't go to E unless we're currently on z
		if letterNext == 'E' && letter != 'z' {
			return
		}
		nextPos := point{y: nextY, x: nextX}
		if !t.visited[nextPos] {
			t.nextPositions[nextPos] = struct{}{}
			t.visited[nextPos] = true
		}
	}
}

func (t *traverser) print() {
	for j, line := range t.lines {
		for i, r := range line {
			if i == 0 {
				fmt.Println()
			}

			p := point{y: j, x: i}
			_, ok := t.positions[p]
			if ok {
				fmt.Print("*")
			} else {
				if t.visited[p] {
					fmt.Print("-")
				} else {
					fmt.Printf("%s", string(r))
				}
			}
		}
	}
	fmt.Println()
}

type debugLogger struct {
	debug bool
}

func (d *debugLogger) Printf(format string, v ...any) {
	if !d.debug {
		return
	}
	log.Printf(format, v...)
}
