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

	//inputs, err := openFile("example.txt")
	inputs, err := openFile("../inputs-day5.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range inputs {
		s := strings.Split(line, " ")
		move, _ := strconv.Atoi(s[1])
		src, _ := strconv.Atoi(s[3])
		dst, _ := strconv.Atoi(s[5])

		// Load
		load := []string{}
		cinta := revertSlice(crane[src])
		for x := 0; x < move; x++ {
			load = append(load, cinta[x])
		}

		// Source
		srcslice := []string{}
		for x := 0; x < (len(crane[src]) - move); x++ {
			srcslice = append(srcslice, crane[src][x])
		}

		// Destination
		dstslice := crane[dst]
		//box := revertSlice(load)
		for x := 0; x < len(load); x++ {
			dstslice = append(dstslice, load[x])
		}

		// Update
		crane[src] = srcslice
		crane[dst] = dstslice
	}
	// Print result
	out := []string{}
	for x := 1; x <= 9; x++ {
		out = append(out, crane[x][(len(crane[x])-1)])
	}
	fmt.Println(out)
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
	for x := (len(array) - 1); x >= 0; x-- {
		result = append(result, array[x])
	}
	return result
}
