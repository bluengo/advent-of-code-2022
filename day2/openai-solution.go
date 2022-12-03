package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Let's play rock, paper, scissors!")

	// create a map of valid moves
	moves := map[string]bool{
		"rock":     true,
		"paper":    true,
		"scissors": true,
	}

	// create a map of the winning moves
	winners := map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	// create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// keep playing the game until the user quits
	for {
		// prompt the first player for their move
		fmt.Println("Player 1, make your move:")
		scanner.Scan()
		p1Move := strings.ToLower(scanner.Text())

		// check if the first player's move is valid
		if !moves[p1Move] {
			fmt.Println("Invalid move!")
			continue
		}

		// prompt the second player for their move
		fmt.Println("Player 2, make your move:")
		scanner.Scan()
		p2Move := strings.ToLower(scanner.Text())

		// check if the second player's move is valid
		if !moves[p2Move] {
			fmt.Println("Invalid move!")
			continue
		}

		// check if the game was a tie
		if p1Move == p2Move {
			fmt.Println("It's a tie!")
			continue
		}

		// determine the winner of the game
		var winner string
		if winners[p1Move] == p2Move {
			winner = "Player 1"
		} else {
			winner = "Player 2"
		}

		// print the result of the game
		fmt.Printf("%s wins!\n", winner)
	}
}
