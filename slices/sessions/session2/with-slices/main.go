package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// dividing into arbitrary sessions {session2}
// Unique random number generator using slices
func main() {
	// create a random seed
	rand.Seed(time.Now().UnixNano())

	// check for the command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments provided")
		return
	}

	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid arguments provided \n")
		return
	}

	// store unique numbers to be generated into a slice
	var uniques []int

	// outer loop
loop:
	for len(uniques) < max {
		n := rand.Intn(max + 1)
		fmt.Print(n, " ")

		// this will not be touched if the uniques is nil
		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques = append(uniques, n)
	}

	fmt.Println("\n\n unique numbers", uniques)
}
