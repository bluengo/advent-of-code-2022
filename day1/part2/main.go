package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elfs []int
	var candy_bag []int

	lines, err := openFile("inputs-day1.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		if num != 0 {
			// If the numer is not 0, append calories to "bag"
			candy_bag = append(candy_bag, num)
		} else {
			// A zero means space in the input file, therefore, a new elf
			// Sum all calories for this Elf
			total := sumCals(candy_bag)
			elfs = append(elfs, total)
			candy_bag = nil
			continue
		}
	}

	fattest := highestCals(elfs)
	fmt.Println(fattest)
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

func sumCals(candy_bag []int) int {
	sum := 0
	for _, candy := range candy_bag {
		sum = sum + candy
	}
	return sum
}

func highestCals(elfs []int) int {
	total := 0
	// Sort elfs from higher to lower
	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))
	//fmt.Println(elfs)
	for x := 0; x < 3; x++ {
		total = total + elfs[x]
	}
	return total
}
