package main

/*
  Day 7. Part 1:
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var folderSizes = map[string]int{
	"/": 0,
}
var sum int

func main() {
	pwd := []string{"/"}
	//numregex := regexp.MustCompile(`^[0-9]$`)
	//charegex := regexp.MustCompile(`^[a-z]$`)

	inputs, err := openFile("../inputs-day7.txt")
	if err != nil {
		panic(err)
	}

	// input into slice of slices
	var lines [][]string
	for _, line := range inputs {
		command := strings.Split(line, " ")
		lines = append(lines, command)
	}

	// read and execute commands
	// skip [0] ("cd /")
	for x := 1; x < len(inputs)-1; x++ {
		stdin := lines[x]
		fmt.Println(stdin)
		// in case of "<size> <file>" I try to find regex
		//newsize := numregex.FindString(stdin[0])
		//filename := charegex.FindAllStringSubmatch(stdin[1], 1)[0]
		switch stdin[0] {
		// Command ======
		case "$":
			switch stdin[1] {
			case "cd":
				// Stop counting previous folder
				saveSum(strings.Join(pwd, " "), sum)
				switch stdin[2] {
				case "..":
					// remove last dir from pwd
					if len(pwd) > 1 {
						pwd = pwd[:len(pwd)-1]
					}
					//DEBUG
					fmt.Println("PWD = ", pwd)
				default:
					folder := fmt.Sprintf("%s", stdin[2])
					pwd = append(pwd, folder)
					fmt.Println("PWD = ", pwd)
				}
			case "ls":
				fmt.Println("ls to ", pwd)
			}
			// ================
		// Dir
		case "dir":
			fmt.Println("Dir: ", stdin)
			// ================
		// File
		default:
			size, err := strconv.Atoi(stdin[0])
			if err != nil {
				panic(err)
			}
			fmt.Printf("File %s with size %d\n", stdin[1], size)
			sum = sum + size
			// ================
		}
	}
	fmt.Println("Finish")
	for k, v := range folderSizes {
		fmt.Printf("The size of %s is: %d\n", k, v)
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

func saveSum(path string, sum int) {
	folderSizes[path] = sum
}
