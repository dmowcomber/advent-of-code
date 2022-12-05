package input

import (
	"bufio"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	result := make([]string, 0)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}
	return result, nil
}
