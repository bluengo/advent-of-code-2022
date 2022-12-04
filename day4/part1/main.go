package main

/*
  Day 3. Part 1:
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	inputs, err := openFile("../part1/inputs-day4.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}

	for _, line := range inputs {
		elves := strings.Split(line, ",")
		elf1 := genRange(elves[0])
		elf2 := genRange(elves[1])
		if isSubset(elf1, elf2) {
			total++
		}
	}
	fmt.Println("Total: ", total)
}

func openFile(file string) ([]string, error) {
	var fileLines []string
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening inputs file")
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}

func genRange(s string) []int {
	var rang []int
	nums := strings.Split(s, "-")
	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])

	for x := min; x <= max; x++ {
		rang = append(rang, x)
	}

	return rang
}

func isSubset(a, b []int) bool {
	switch {
	case a[0] == b[0]:
		return true
	case a[0] < b[0]:
		if a[len(a)-1] < b[len(b)-1] {
			return false
		} else {
			return true
		}
	case a[0] > b[0]:
		if a[len(a)-1] > b[len(b)-1] {
			return false
		} else {
			return true
		}
	}
	return false
}
