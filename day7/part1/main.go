package main

/*
  Day 6. Part 1:
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputs, err := openFile("../inputs-day7.txt")
	if err != nil {
		panic(err)
	}

	var lines [][]string
	for _, line := range inputs {
		command := strings.Split(line, " ")
		lines = append(lines, command)
	}
	fmt.Println(lines)
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
