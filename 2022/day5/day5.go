package main

import (
	"log"
	"regexp"
	"strconv"

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
		log.Printf("line: %s", line)
		firstRune := line[1]
		_, err := strconv.Atoi(string(firstRune))
		if err == nil {
			// end of crates
			break
		}
		crateLines = append(crateLines, line)
		endOfCrates++
	}

	crates := getCrates(crateLines)
	log.Printf("crates: %#v", crates)

	endOfCrates += 2
	log.Printf("start of moves line: %s", lines[endOfCrates])

	for i := endOfCrates; i < len(lines); i++ {
		line := lines[i]
		count, from, to := getMove(line)

		for i := 0; i < count; i++ {
			x, crateFrom := stackPop(crates[from])
			crates[from] = crateFrom

			// add to crateTo
			log.Printf("moving %s", x)
			crateTo := append([]string{x}, crates[to]...)
			// log.Printf("new crateTo: %v", crateTo)
			crates[to] = crateTo
		}
	}
	log.Printf("crates: %#v", crates)

	answer := ""
	for i := 1; i < len(crates)+1; i++ {
		crate := crates[i]
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

func getCrates(lines []string) map[int][]string {
	crates := make(map[int][]string)
	for _, line := range lines {
		index := 1
		for i := 1; i < len(line); i += 4 {
			r := line[i]
			if string(r) != " " {
				crate, ok := crates[index]
				if !ok {
					crate = make([]string, 0)
				}

				crate = append(crate, string(r))
				crates[index] = crate
			}
			index++
		}
	}
	return crates

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
		log.Printf("line: %s", line)
		firstRune := line[1]
		_, err := strconv.Atoi(string(firstRune))
		if err == nil {
			// end of crates
			break
		}
		crateLines = append(crateLines, line)
		endOfCrates++
	}

	crates := getCrates(crateLines)
	log.Printf("crates: %#v", crates)

	endOfCrates += 2
	log.Printf("start of moves line: %s", lines[endOfCrates])

	for i := endOfCrates; i < len(lines); i++ {
		line := lines[i]
		count, from, to := getMove(line)

		movingStack := make([]string, 0)
		for i := 0; i < count; i++ {
			x, crateFrom := stackPop(crates[from])
			crates[from] = crateFrom

			movingStack = append(movingStack, x)
		}
		crates[to] = append(movingStack, crates[to]...)

	}
	log.Printf("crates: %#v", crates)

	answer := ""
	for i := 1; i < len(crates)+1; i++ {
		crate := crates[i]
		// log.Printf("%d crate: %#v", i, crate)
		_ = i
		x := crate[0]
		answer = answer + x
	}

	return answer, nil
}
