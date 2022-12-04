package main

func getCommonLetter(lines []string) rune {
	commonLetters := make(map[rune]struct{})
	for _, lineRune := range lines[0] {
		commonLetters[lineRune] = struct{}{}
	}
	for _, line := range lines {
		newCommonLetters := make(map[rune]struct{})
		for _, lineRune := range line {
			_, ok := commonLetters[lineRune]
			if ok {
				newCommonLetters[lineRune] = struct{}{}
			}
		}
		commonLetters = newCommonLetters
		// log.Printf("current common letters: %#v", commonLetters)
	}
	// log.Printf("found common letters: %#v", commonLetters)

	for r := range commonLetters {
		return r // return the first rune
	}
	return 0
}

func getLetterPriority(c rune) int {
	letterPrio := 0
	if int(c) >= 97 { // a
		letterPrio = int(c) - 97 + 1
	} else if int(c) <= 90 { // a
		// 65 - 90 / A - Z
		letterPrio = int(c) - 65 + 27
	}
	return letterPrio
}

func getPrirorityPart2(filename string) (int, error) {
	lines, err := readLines(filename)
	if err != nil {
		return 0, err
	}

	var totalPriority int
	for i := 0; i < len(lines); i += 3 {
		linesGroup1 := lines[i : i+3]
		// log.Printf("linesGroup1: %+v", linesGroup1)
		c1 := getCommonLetter(linesGroup1)
		// log.Printf("found letters %v", string(c1))
		totalPriority += getLetterPriority(c1)
	}

	return totalPriority, nil
}
