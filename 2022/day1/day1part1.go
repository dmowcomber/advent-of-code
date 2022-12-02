package main

import (
	"bufio"
	"os"
	"strconv"
)

func getMaxCaloies() (int, error) {
	lines, err := readCalorieLines()
	if err != nil {
		return 0, err
	}

	var caloriesMax, caloriesCount int

	for _, line := range lines {
		if line == "" {
			// handle end of elf
			if caloriesCount > caloriesMax {
				caloriesMax = caloriesCount
			}
			caloriesCount = 0
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		caloriesCount += i
	}
	return caloriesMax, nil
}

func readCalorieLines() ([]string, error) {
	readFile, err := os.Open("./elf-calories.txt")
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
