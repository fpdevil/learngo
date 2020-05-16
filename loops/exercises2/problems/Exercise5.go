package problems

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

// Exercise5 - Enough Picks
func Exercise5(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/05-enough-picks Exercise: Enough Picks ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Println(usage, maxTurns)
		return
	}

	var maxval = 10

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a number", args[0])
		return
	}

	t := time.Now()
	rand.Seed(t.UnixNano())

	if maxval < guess {
		maxval = guess
	}

	for i := 0; i < maxval; i++ {
		n := rand.Intn(guess + 1)
		// fmt.Printf("turn value %d\n", n)
		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}
	fmt.Println("☠️  YOU LOST!")
}
