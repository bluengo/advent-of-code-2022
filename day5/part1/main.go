package main

/*
  Day 3. Part 1:
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var crane1 = []string{"D", "L", "V", "T", "M", "H", "F"}
	var crane2 = []string{"H", "Q", "G", "J", "C", "T", "N", "P"}
	var crane3 = []string{"R", "S", "D", "M", "P", "H"}
	var crane4 = []string{"L", "B", "V", "F"}
	var crane5 = []string{"N", "H", "G", "L", "Q"}
	var crane6 = []string{"W", "B", "D", "G", "R", "M", "P"}
	var crane7 = []string{"G", "M", "N", "R", "C", "H", "L", "Q"}
	var crane8 = []string{"C", "L", "W"}
	var crane9 = []string{"R", "D", "L", "Q", "J", "Z", "M", "T"}
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
