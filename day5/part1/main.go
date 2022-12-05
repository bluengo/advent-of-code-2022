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
	var crane1 []string
	var crane2 []string
	var crane3 []string
	var crane4 []string
	var crane5 []string
	var crane6 []string
	var crane7 []string
	var crane8 []string
	var crane9 []string
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
