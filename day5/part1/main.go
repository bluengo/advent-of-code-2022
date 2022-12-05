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
	var crane = [][]string{
		{"NULL"}, // To start from index 1:
		{"D", "L", "V", "T", "M", "H", "F"},
		{"H", "Q", "G", "J", "C", "T", "N", "P"},
		{"R", "S", "D", "M", "P", "H"},
		{"L", "B", "V", "F"},
		{"N", "H", "G", "L", "Q"},
		{"W", "B", "D", "G", "R", "M", "P"},
		{"G", "M", "N", "R", "C", "H", "L", "Q"},
		{"C", "L", "W"},
		{"R", "D", "L", "Q", "J", "Z", "M", "T"},
	}

	inputs, err := openFile("example.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range inputs {
		fmt.Println(line)
		s := strings.Split(line, " ")
		move, _ := strconv.Atoi(s[1])
		src, _ := strconv.Atoi(s[3])
		dst, _ := strconv.Atoi(s[5])
		//fmt.Println(move, src, dst)

		cargo := revertSlice(crane[src][:move])

		// Move cargo:
		for x := 0; x < move; x++ {
			crane[dst] = append(crane[dst], cargo[x])
		}
		fmt.Println(crane)
	}
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

func revertSlice(array []string) []string {
	var result []string
	for _, x := range array {
		result = append(result, x)
	}
	return result
}
