package problems

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

// Exercise6 - Dynamic Difficulty
func Exercise6(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/06-dynamic-difficulty Exercise: Dynamic Difficulty	----")

	args := strings.Split(vals, ",")
	fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Println(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q is not a number", args[0])
		return
	}

	t := time.Now()
	rand.Seed(t.UnixNano())

	var difficulty int

	if guess <= maxTurns {
		difficulty = guess
	} else {
		difficulty = maxTurns + guess/4
	}

	// fmt.Println("chosen difficulty ", difficulty)
	for turn := 0; turn < difficulty; turn++ {
		n := rand.Intn(guess + 1)
		// fmt.Printf("guessed value %d\n", n)
		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}
	fmt.Println("☠️  YOU LOST!")
}
