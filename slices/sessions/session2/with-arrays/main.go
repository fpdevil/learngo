package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Unique random number generator
func main() {
	rand.Seed(time.Now().UnixNano())

	// generate max number of random numbers
	const max = 5
	// with arrays, we cabnot pass the value of max
	// dynamically as under
	// max, _ := strconv.Atoi(os.Args[1])

	// store unique numbers into an array
	var uniques [max]int

loop:
	for found := 0; found < max; {
		n := rand.Intn(max + 1)
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques[found] = n
		// this ensures that increment happens only if
		// a unique number is found
		found++
	}

	fmt.Println("\n\n unique numbers", uniques)
}
