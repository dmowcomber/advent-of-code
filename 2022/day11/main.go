package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/dmowcomber/advent-of-code/input"
)

type monkey struct {
	items       []int
	operator    string
	operand     string
	mod         int
	trueMonkey  int
	falseMonkey int
}

func getMonkeyBusinessLevel(filename string, rounds int, relax bool) (int, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	monkeyIndex := -1

	monkeys := make(map[int]*monkey)
	for _, line := range lines {
		if line == "" {
			monkeyIndex = -1
			continue
		}

		// log.Printf("line: %s", line)
		if strings.HasPrefix(line, "Monkey") {
			monkeyIndex, _ = strconv.Atoi(line[7:8])
			monkeys[monkeyIndex] = &monkey{}
			continue
		}

		if strings.Contains(line, "Starting items: ") {
			log.Printf("items: %s", line[18:])
			itemsStrs := strings.Split(line[18:], ", ")
			items := make([]int, 0, len(itemsStrs))
			for _, itemStr := range itemsStrs {
				num, _ := strconv.Atoi(itemStr)
				items = append(items, num)
			}
			monkeys[monkeyIndex].items = items
			continue
		}

		if strings.Contains(line, "Operation:") {
			operator := line[23:24]
			monkeys[monkeyIndex].operator = operator
			monkeys[monkeyIndex].operand = line[25:]
			continue
		}

		if strings.Contains(line, "Test: divisible by ") {
			mod := line[21:]
			monkeys[monkeyIndex].mod, _ = strconv.Atoi(mod)
		}

		if strings.Contains(line, "If true: throw to monkey") {
			trueMonkey := line[29:]
			monkeys[monkeyIndex].trueMonkey, _ = strconv.Atoi(trueMonkey)
		}
		if strings.Contains(line, "If false: throw to monkey") {
			falseMonkey := line[30:]
			monkeys[monkeyIndex].falseMonkey, _ = strconv.Atoi(falseMonkey)
		}

		if line == "" {
			monkeyIndex = -1
			continue
		}
	}
	// log.Printf("monkeys: %#v", monkeys)

	combinedMod := 1
	for _, monk := range monkeys {
		combinedMod *= monk.mod
	}

	monkeyCounts := make(map[int]int)
	for round := 1; round < rounds+1; round++ {
		for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {
			monk := monkeys[monkeyIndex]
			monkeyCounts[monkeyIndex] += len(monk.items)

			// log.Printf("monkey start items: %v", monk.items)
			for itemIndex := range monk.items {
				// log.Printf("monkey start items[i]: %d", monk.items[itemIndex])

				operand := monk.items[itemIndex]
				if monk.operand != "old" {
					operand, _ = strconv.Atoi(monk.operand)
				}

				if monk.operator == "+" {
					monk.items[itemIndex] += operand
				}
				if monk.operator == "*" {
					monk.items[itemIndex] *= operand
				}

				if relax {
					monk.items[itemIndex] /= 3
				} else {
					// keep the end result counts the same but reduce the item values
					monk.items[itemIndex] %= combinedMod
				}

				passMonkeyIndex := monk.falseMonkey
				if monk.items[itemIndex]%monk.mod == 0 {
					passMonkeyIndex = monk.trueMonkey
				}
				monkeys[passMonkeyIndex].items = append(monkeys[passMonkeyIndex].items, monk.items[itemIndex])

			}

			monk.items = nil

		}

		if round == 1 || round == 20 || (round)%1000 == 0 {
			log.Printf("%d - monkeyCounts: %#v", round, monkeyCounts)
		}
	}

	// for monkeyIndex, monk := range monkeys {
	// log.Printf("monkey %d has %v", monkeyIndex, monk.items)
	// }

	log.Printf("monkeyCounts: %#v", monkeyCounts)

	counts := make([]int, 0, len(monkeyCounts))
	for _, count := range monkeyCounts {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	log.Printf("counts: %v", counts)
	return counts[0] * counts[1], nil
}
