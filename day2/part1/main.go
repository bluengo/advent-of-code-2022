package main

/*
  Day 2. Part 1:
  ==============
    rock (1) = A, X
    paper (2) = B, Y
    scissors (3) = C, Z

    loose = 0
    draw = 3
    win = 6

    #Example:

    A Y = 2 + 6 = 8
    B X = 1 + 0 = 1
    C Z = 3 + 3 = 6
                 15
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs, err := openFile("inputs-day2.txt")
	if err != nil {
		panic("Unable to open inputs file")
	}

	points := 0
	for _, match := range inputs {
		players := strings.Split(match, " ")

		switch players[0] {
		case "A":
			switch players[1] {
			case "X":
				points = points + 1 + 3
			case "Y":
				points = points + 2 + 6
			case "Z":
				points = points + 3
			default:
				msg, _ := fmt.Printf("Unknown option %s", players[1])
				panic(msg)
			}
		case "B":
			switch players[1] {
			case "X":
				points = points + 1
			case "Y":
				points = points + 2 + 3
			case "Z":
				points = points + 3 + 6
			default:
				msg, _ := fmt.Printf("Unknown option %s", players[1])
				panic(msg)
			}
		case "C":
			switch players[1] {
			case "X":
				points = points + 1 + 6
			case "Y":
				points = points + 2
			case "Z":
				points = points + 3 + 3
			default:
				msg, _ := fmt.Printf("Unknown option %s", players[1])
				panic(msg)
			}
		default:
			msg, _ := fmt.Printf("Unknown option %s", players[0])
			panic(msg)
		}
	}

	fmt.Printf("Total Score: %d\n", points)
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
