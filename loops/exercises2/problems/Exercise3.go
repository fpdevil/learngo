package problems

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

// Exercise3 - Double Guesses
func Exercise3(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/03-double-guesses Exercise: Double Guesses ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	t := time.Now()
	rand.Seed(t.UnixNano())

	var guess1, guess2 int

	guess1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%q Not a number!\n", args[0])
		return
	}

	if len(args) == 2 {
		guess2, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("%q Not a number!\n", args[1])
			return
		}
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Guess values cannot be negative")
		return
	}

	var maxval int
	if guess2 > guess1 {
		maxval = guess2
	} else {
		maxval = guess1
	}

	for turns := 0; turns <= maxTurns; turns++ {
		n := rand.Intn(maxval + 1)
		// fmt.Printf("system: %d\n", n)
		if n == guess1 || n == guess2 {
			fmt.Println("🎉  YOU WON!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST!")
}
