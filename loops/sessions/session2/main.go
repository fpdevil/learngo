package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Building a Lucky Number Game

const (
	maxTurns = 5 // magic number - less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	// random numbers
	t := time.Now()
	rand.Seed(t.UnixNano())

	// define a guess
	// guess := 10

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number!")
		return
	}

	if guess < 0 {
		fmt.Println("Guess cannot be negative")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
