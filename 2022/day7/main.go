package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dmowcomber/advent-of-code/input"
)

func getPart1(filename string) (int, error) {
	rootDirectory, err := getRootDirectory(filename)
	if err != nil {
		return 0, err
	}

	// sum up all directories within the threshold starting with the root
	sum := 0
	max := 100000
	sum = getSum(sum, max, rootDirectory)
	return sum, nil
}

func getPart2(filename string) (int, error) {
	rootDirectory, err := getRootDirectory(filename)
	if err != nil {
		return 0, err
	}

	deleteable := rootDirectory.Size()
	freeSpace := 70000000 - deleteable
	neededSpace := 30000000 - freeSpace
	deleteable = getDeletable(rootDirectory, deleteable, neededSpace)
	return deleteable, nil
}

func getRootDirectory(filename string) (*directory, error) {
	lines, err := input.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	lines = lines[1:] // ignore the first line which is always `$ cd /`

	current := ""
	fileSizes := make(map[string]int)
	for _, line := range lines {
		num := getFileSize(line)
		if num == "" {
			// handle change directories
			if line == "$ cd .." {
				current = getDirecotryOfFile(current)
				continue
			}
			if strings.HasPrefix(line, "$ cd ") {
				current = current + "/" + strings.TrimPrefix(line, "$ cd ")
				continue
			}
			continue
		}

		// handle file sizes
		i, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}

		// add num and current/filename to maps only.. no dirs
		filename := getFilename(line)
		filePath := current + "/" + filename
		fileSizes[filePath] = i
	}

	// get sizes of immediate directories
	dirSizes := make(map[string]int)
	dirs := make(map[string]*directory)
	for filePath, fileSize := range fileSizes {
		dir := getDirecotryOfFile(filePath)
		dirSizes[dir] = dirSizes[dir] + fileSize
		if dirs[dir] == nil {
			dirs[dir] = &directory{
				childDirs: make(map[string]*directory),
			}
		}
		dirs[dir].fileSizes = dirSizes[dir]
		linkToParent(dir, dirs)
	}

	// return the root directory which has all the nested directories linked
	return dirs[""], nil
}

func getSum(sum, max int, directry *directory) int {
	size := directry.Size()
	if size <= max {
		sum += size
	}
	for _, childDirectory := range directry.childDirs {
		sum = getSum(sum, max, childDirectory)
	}
	return sum
}

func getDeletable(dir *directory, deletable, neededSpace int) int {
	if dir.Size() >= neededSpace && dir.Size() <= deletable {
		deletable = dir.Size()
	}
	for _, childDir := range dir.childDirs {
		deletable = getDeletable(childDir, deletable, neededSpace)
	}
	return deletable
}

func getFileSize(input string) string {
	r := regexp.MustCompile(`^(\d+) .*$`)
	matches := r.FindStringSubmatch(input)
	if len(matches) != len(r.SubexpNames()) {
		return ""
	}
	return matches[1]
}

func getFilename(input string) string {
	r := regexp.MustCompile(`^\d+ (.+)$`)
	matches := r.FindStringSubmatch(input)
	if len(matches) != len(r.SubexpNames()) {
		return ""
	}
	return matches[1]
}

func linkToParent(dir string, dirs map[string]*directory) {
	// root dir has no parents
	if dir == "" {
		return
	}

	parentDir := getDirecotryOfFile(dir)
	if dirs[parentDir] == nil {
		dirs[parentDir] = &directory{
			childDirs: make(map[string]*directory),
		}
	}
	parentDirectory := dirs[parentDir]
	parentDirectory.childDirs[dir] = dirs[dir]
	dirs[parentDir] = parentDirectory

	// recurse to children of children
	linkToParent(parentDir, dirs)
}

// getDirecotryOfFile gets the directory of a file (or directory)
func getDirecotryOfFile(filePath string) string {
	fileParts := strings.Split(filePath, "/")
	if len(fileParts) < 1 {
		return "" // root dir
	}
	dirParts := fileParts[:len(fileParts)-1]
	dir := strings.Join(dirParts, "/")
	return dir
}

type directory struct {
	fileSizes       int // size of files in the immediate directory
	totalNestedSize int
	childDirs       map[string]*directory
}

func (d *directory) Size() int {
	if d.totalNestedSize != 0 {
		return d.totalNestedSize
	}

	sum := d.fileSizes
	for _, childDirectory := range d.childDirs {
		sum += childDirectory.Size()
	}
	d.totalNestedSize = sum
	return sum
}
