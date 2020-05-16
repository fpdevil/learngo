package problems

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

// using the same values from the original video
// session shown as these are just constants
const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.

Wanna play?
`
)

// Exercise1 - First Turn Winner
func Exercise1(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/08-lucky-number-exercises Exercise: First Turn Winner ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" {
		fmt.Println("Guess a number.")
		return
	}

	t := time.Now()
	rand.Seed(t.UnixNano())

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number!")
		return
	}

	if guess < 0 {
		fmt.Println("Guess cannot be negative")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!!")
			} else {
				fmt.Println("🎉  YOU JUST WON!")
			}
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Better Luck Next Time!")
}
