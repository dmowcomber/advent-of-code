package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/dmowcomber/advent-of-code/input"
)

func getPart1(filename string) (string, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return "", err
	}

	crateLines := make([]string, 0)

	endOfCrates := 0
	for _, line := range lines {
		log.Printf("input line: %s", line)
		firstRune := line[1]
		_, err := strconv.Atoi(string(firstRune))
		if err == nil {
			// end of crates
			break
		}
		crateLines = append(crateLines, line)
		endOfCrates++
	}

	crates := getCrateStacks(crateLines)
	log.Printf("stacks:\n%s\n", crates.String())

	endOfCrates += 2
	log.Printf("start of moves line: %s", lines[endOfCrates])

	for i := endOfCrates; i < len(lines); i++ {
		line := lines[i]
		count, from, to := getMove(line)

		for i := 0; i < count; i++ {
			x, crateFrom := stackPop(crates.stacks[from])
			crates.stacks[from] = crateFrom

			// add to crateTo
			// log.Printf("moving %s from %d to %d", x, from, to)
			crateTo := append([]string{x}, crates.stacks[to]...)
			// log.Printf("new crateTo: %v", crateTo)
			crates.stacks[to] = crateTo
			// log.Printf("stacks:\n%s\n", crates.String())
		}
	}

	log.Printf("stacks:\n%s\n", crates.String())

	answer := ""
	for i := 1; i < len(crates.stacks)+1; i++ {
		crate := crates.stacks[i]
		// log.Printf("%d crate: %#v", i, crate)
		_ = i
		x := crate[0]
		answer = answer + x
	}

	return answer, nil
}

func stackPop(crate []string) (string, []string) {
	x, crate := crate[0], crate[1:]
	// log.Printf("new crate after pop: %v", crate)
	return x, crate
}

func getCrateStacks(lines []string) crateStacks {
	stacks := make(map[int][]string)
	for _, line := range lines {
		index := 1
		for i := 1; i < len(line); i += 4 {
			r := line[i]
			if string(r) != " " {
				crate, ok := stacks[index]
				if !ok {
					crate = make([]string, 0)
				}

				crate = append(crate, string(r))
				stacks[index] = crate
			}
			index++
		}
	}

	return crateStacks{stacks: stacks}

	//[x] [y] [z]
	//0123456789

	// [D]
	// [N] [C]
	// [Z] [M] [P]
	//  1   2   3

	// move 1 from 2 to 1
	// move 3 from 1 to 3
	// move 2 from 2 to 1
	// move 1 from 1 to 2
}

func getMove(line string) (count, from, to int) {
	// Example:
	// move 1 from 2 to 1
	// move 3 from 1 to 3
	// move 2 from 2 to 1
	// move 1 from 1 to 2
	re := regexp.MustCompile(`move (?P<count>\d+) from (?P<from>\d+) to (?P<to>\d+)`)
	matches := re.FindStringSubmatch(line)
	countIndex := re.SubexpIndex("count")
	countStr := matches[countIndex]

	fromIndex := re.SubexpIndex("from")
	fromStr := matches[fromIndex]

	toIndex := re.SubexpIndex("to")
	toStr := matches[toIndex]

	// fmt.Printf("m %s f %s t %s\n", countStr, fromStr, toStr)
	count, _ = strconv.Atoi(countStr)
	from, _ = strconv.Atoi(fromStr)
	to, _ = strconv.Atoi(toStr)
	// fmt.Printf("m %d f %d t %d\n", count, from, to)
	return
}

func getPart2(filename string) (string, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return "", err
	}

	crateLines := make([]string, 0)

	endOfCrates := 0
	for _, line := range lines {
		log.Printf("input line: %s", line)
		firstRune := line[1]
		_, err := strconv.Atoi(string(firstRune))
		if err == nil {
			// end of crates
			break
		}
		crateLines = append(crateLines, line)
		endOfCrates++
	}

	crates := getCrateStacks(crateLines)
	log.Printf("starting stacks:\n%s\n", crates.String())

	// log.Printf("crates: %#v", crates.stacks)

	endOfCrates += 2
	log.Printf("start of moves line: %s", lines[endOfCrates])

	for i := endOfCrates; i < len(lines); i++ {
		line := lines[i]
		count, from, to := getMove(line)

		movingStack := make([]string, 0)
		for i := 0; i < count; i++ {
			x, crateFrom := stackPop(crates.stacks[from])
			crates.stacks[from] = crateFrom

			movingStack = append(movingStack, x)
		}
		crates.stacks[to] = append(movingStack, crates.stacks[to]...)

	}
	// log.Printf("crates: %#v", crates)
	log.Printf("result stacks:\n%s\n", crates.String())

	answer := ""
	// retain order by looping crate stack indicies
	for i := 1; i < len(crates.stacks)+1; i++ {
		crate := crates.stacks[i]
		// log.Printf("%d crate: %#v", i, crate)
		_ = i
		x := crate[0]
		answer = answer + x
	}

	return answer, nil
}

type crateStacks struct {
	stacks map[int][]string
}

// for fun return a string representation of the crate stacks
func (c *crateStacks) String() string {
	// sideways stack visualization:
	// abc
	// d
	// ef

	// Desired stack visualization:
	// a
	// b   e
	// c d f

	// copy stacks so we can pop items off the stacks when constructing the resulting string
	tmpStacks := make(map[int][]string)
	for i, stack := range c.stacks {
		tmpStack := make([]string, len(stack))
		copy(tmpStack, stack)
		tmpStacks[i] = tmpStack
	}

	result := ""

	// rotate:
	for heightIndex := 0; ; heightIndex++ {
		allEmpty := true
		resultLine := ""
		for i := len(tmpStacks); i > 0; i-- {
			// prepending
			// starting with the bottom
			// since we don't know how tall each stack will be:

			// check size
			if len(tmpStacks[i]) == 0 {
				// append empty string?
				resultLine = "    " + resultLine
				continue
			}
			allEmpty = false

			// get last item
			// log.Printf("%d - tmpStacks[i]: %s", i, tmpStacks[i])
			x := tmpStacks[i][len(tmpStacks[i])-1]
			tmpStacks[i] = tmpStacks[i][:len(tmpStacks[i])-1]
			resultLine = fmt.Sprintf("[%s] ", x) + resultLine
		}

		if allEmpty {
			result = strings.TrimLeft(result, "\n")
			break
		}
		result = resultLine + result
		result = "\n" + result
	}
	return result
}
