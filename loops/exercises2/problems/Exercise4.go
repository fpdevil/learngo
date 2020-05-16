package problems

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	usage4 = `
	usage: go run main.go [guess] or
		   go run main.go -v [guess]
	`
)

// Exercise4 - Verbose Mode
func Exercise4(vals string) {
	fmt.Println("")
	fmt.Println("---- 13-loops/04-verbose-mode Exercise: Verbose Mode ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		fmt.Println(usage4)
		return
	}

	var guess int
	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Printf("%q is not a number", args[len(args)-1])
		return
	}
	// fmt.Println(strings.Index(vals, "-v"))

	t := time.Now()
	rand.Seed(t.UnixNano())

	if guess < 0 {
		fmt.Println("Guess value cannot be negative")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if verbose {
			fmt.Printf("%d ", n)
		}
		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}
	fmt.Println("☠️  YOU LOST!")
}
