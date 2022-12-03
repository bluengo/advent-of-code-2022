package main

/*
  Day 3. Part 1:
  ==============
    Example:
    vJrwpWtwJgWrhcsFMMfFFhFp			p
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL	L
	PmmdzqPrVvPwwTWBwg					P
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn		v
	ttgJtRGJQctTZtZT					t
	CrZsJsPPZsGzwwsLwLmpwMDw			s

	a-z = 1-26
	A-Z = 27-52
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs, err := openFile("inputs-day3.txt")
	//inputs, err := openFile("example.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}

	var common []string
	var sum int

	for _, line := range inputs {
		rucksack := strings.Split(line, "")
		half := len(rucksack) / 2
		comp1 := rucksack[:half]
		comp2 := rucksack[half:]
		//fmt.Println(comp1, comp2)
		for _, x := range comp1 {
			if checkIfIsInSlice(x, comp2) {
				common = append(common, x)
				break
			}
		}
	}
	alfa := generateAlphabet()
	//fmt.Println(alfa, len(alfa))
	//fmt.Println(common)
	for _, x := range common {
		index, err := getIndex(x, alfa)
		if err != nil {
			panic(err)
		}
		sum = sum + index
	}
	fmt.Println(sum)
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

func checkIfIsInSlice(a string, b []string) bool {
	for _, item := range b {
		if a == item {
			return true
		}
	}
	return false
}

func generateAlphabet() []string {
	alphabet := make([]string, 0, 52)
	var ch byte

	// lowercase
	for ch = 'a'; ch <= 'z'; ch++ {
		chs := string(ch)
		alphabet = append(alphabet, chs)
	}

	// uppsercase
	for ch = 'A'; ch <= 'Z'; ch++ {
		chs := string(ch)
		alphabet = append(alphabet, chs)
	}
	return alphabet
}

func getIndex(c string, arr []string) (int, error) {
	for x := 0; x <= len(arr); x++ {
		if c == arr[x] {
			return x + 1, nil
		}
	}
	return 0, errors.New("Character not found in array")
}
