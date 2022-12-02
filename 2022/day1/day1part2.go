package main

import (
	"errors"
	"log"
	"sort"
	"strconv"
)

func getSumTop3Caloies() (int, error) {
	lines, err := readCalorieLines()
	if err != nil {
		return 0, err
	}

	calories := make([]int, 0, len(lines))

	var caloriesCount int
	for _, line := range lines {
		if line == "" {
			// handle end of elf
			calories = append(calories, caloriesCount)
			caloriesCount = 0
			continue
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		caloriesCount += i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	log.Printf("calories: %#v", calories)

	if len(calories) < 3 {
		return 0, errors.New("unexpected size of calories slice")
	}

	sumTop3calories := calories[0] + calories[1] + calories[2]
	return sumTop3calories, nil
}
