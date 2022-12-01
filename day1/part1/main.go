package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var elfs []int
	//var candy_bag []int

	lines, err := openFile("inputs-day1.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		elfs = append(elfs, num)
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

func highestCals(elfs []int) int {
	highest := 0
	for _, elf := range elfs {
		if elf > highest {
			highest = elf
		}
	}
	return highest
}

//func main() {
//  myvar := &[]int{3, 4, 5}
//  fmt.Println(&myvar)
//  fmt.Println((*myvar)[1])
//  for i, n := range *myvar {
//    fmt.Printf("Item %d: %d\n", i, n)
//  }
//}
