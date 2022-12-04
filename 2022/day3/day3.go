package main

import (
	"bufio"
	"os"
)

func getPrirority(filename string) (int, error) {
	lines, err := readLines(filename)
	if err != nil {
		return 0, err
	}

	var prioTotal int
	for _, line := range lines {
		// log.Println("\n\nnew line")
		// if len(line) %2 == 0 {

		// }
		length := len(line)

		middle := length / 2

		part1 := line[0:middle]
		part2 := line[(middle):]
		// log.Printf("part1: %s, part2: %s", part1, part2)
		letters := make(map[rune]struct{})

		for _, c := range part1 {
			letters[c] = struct{}{}
		}

		for _, c := range part2 {
			_, ok := letters[c]
			if ok {
				// calculate/add prio

				letterPrio := 0
				if int(c) >= 97 { // a
					letterPrio = int(c) - 97 + 1
				} else if int(c) <= 90 { // a
					// 65 - 90 / A - Z
					letterPrio = int(c) - 65 + 27
				}
				// log.Printf("found letter in both: %s, %d", string(c), letterPrio)

				prioTotal += letterPrio
				goto LOOP
			} else {
				continue
			}
		}
	LOOP:
		continue
	}

	return prioTotal, nil
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
