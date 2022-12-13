package main

import (
	"regexp"

	"github.com/dmowcomber/advent-of-code/input"
)

func getPart1(filename string) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	for _, line := range lines {
		_ = line

	}

	return 0, nil
}

func getPart2(filename string) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	for _, line := range lines {
		_ = line

	}

	return 0, nil
}

func getRegexExample(input string) string {
	r := regexp.MustCompile(`^(\d+) .*`)
	matches := r.FindStringSubmatch(input)
	if len(matches) != len(r.SubexpNames()) {
		return ""
	}

	return matches[1]
}
