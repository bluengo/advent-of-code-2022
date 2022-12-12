package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	number         int
	starting_items []int
	operation      string
	test           string
	if_true        string
	if_false       string
}

func main() {
	inputfile, err := openFile("../inputs-day11.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}
	//fmt.Println(inputfile[0])
	monetes, err := parseInput(inputfile)
	if err != nil {
		panic(err)
	}
	fmt.Println(monetes)
}

// Functions:
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

func parseInput(input []string) (map[int]Monkey, error) {
	var monkeys = make(map[int]Monkey)
	for _, l := range input {
		line := strings.Fields(l)

		if len(line) != 0 {
			newmonkey := Monkey{}
			switch line[0] {
			case "Monkey":
				num := strings.Split(line[1], "")
				monkey_number, err := strconv.Atoi(num[0])
				if err != nil {
					return nil, err
				}
				newmonkey.number = monkey_number
				monkeys[monkey_number] = newmonkey
			case "Starting":
				var items []int
				for _, nstr := range line[2:] {
					nint, _ := strconv.Atoi(nstr)
					items = append(items, nint)
				}
				for _, nstr := range line[2:] {
					n := strings.Split(nstr, "")[:2]
					num := strings.Join(n, "")
					nint, err := strconv.Atoi(num)
					if err != nil {
						return nil, err
					}
					items = append(items, nint)
				}
				fmt.Println(items)
				newmonkey.starting_items = items
			default:
				fmt.Println(monkeys)
			}
		}
	}
	return monkeys, nil
}
