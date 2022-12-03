package main

/*
  Day 3. Part 2:
  ==============
    Example:
    vJrwpWtwJgWrhcsFMMfFFhFp			r
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL	r
	PmmdzqPrVvPwwTWBwg					r
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn		Z
	ttgJtRGJQctTZtZT					Z
	CrZsJsPPZsGzwwsLwLmpwMDw			Z
									   50
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
	inputs, err := openFile("../part1/inputs-day3.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}

	var group []string
	alfa := generateAlphabet()
	var sum int

	for x := 0; x <= len(inputs)-1; x = x + 3 {
		group = append(group, inputs[x])
		group = append(group, inputs[x+1])
		group = append(group, inputs[x+2])

		value, err := getIndex(getBadge(group), alfa)
		if err != nil {
			panic(err)
		}

		sum = sum + value
		group = []string{}
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

func getBadge(group []string) string {
	elf1 := strings.Split(group[0], "")
	elf2 := strings.Split(group[1], "")
	elf3 := strings.Split(group[2], "")

	for _, char1 := range elf1 {
		for _, char2 := range elf2 {
			if char1 == char2 {
				for _, char3 := range elf3 {
					if char2 == char3 {
						return char1
					}
				}
			}
		}
	}
	return "Letter not found"
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
