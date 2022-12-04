package main

func getCountPart2(filename string) (int, error) {
	lines, err := readLines(filename)
	if err != nil {
		return 0, err
	}

	var count int
	for _, line := range lines {
		a, b, x, y, err := getNums(line)
		if err != nil {
			return 0, nil
		}
		// log.Printf("a,b,x,y: %v %v %v %v", a, b, x, y)

		// 1-5,4-7
		// a-b,x-y
		// 1<=4,5>=4
		// 5<=6,6<=6

		// 6-6,4-6
		// a<=y,b>=y

		// 4-6,6-6
		// a<=6,6>=6

		// non-subset examples
		if a <= x && b >= x {
			count++
		} else if a <= y && b >= y {
			count++

			// subset 5-6,2-7:
		} else if a <= x && y <= b {
			count++
		} else if x <= a && b <= y {
			count++
		}
	}

	return count, nil
}
