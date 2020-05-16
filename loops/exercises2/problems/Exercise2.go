package problems

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

const (
	won1  = "YOU WON"
	won2  = "YOU'RE AWESOME"
	lose1 = "LOSER!"
	lose2 = "YOU LOST. TRY AGAIN?"
)

// Exercise2 - Random Messages
func Exercise2(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/02-random-messages Exercise: Sum the Numbers: Random Messages ----")

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

	for turn := 1; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			switch rand.Intn(2) {
			case 0:
				fmt.Println(won1)
			case 1:
				fmt.Println(won2)
			}
			return
		}
	}

	switch rand.Intn(2) {
	case 0:
		fmt.Println(lose1)
	case 1:
		fmt.Println(lose2)
	}
}
