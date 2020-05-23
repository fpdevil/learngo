package main

import (
	"fmt"
	// s "github.com/inancgumus/prettyslice"
	// "github.com/inancgumus/screen"
)

func main() {
	fmt.Println("Mechanics of append continued for larger values...")
	fmt.Println("Growth of Backing Array in Animation")

	// oldCap stores previous capacity for calculating growth ratio
	ages, oldCap := []int{1}, 1.

	// loop for length less than a value
	for len(ages) < 5e4 {
		ages = append(ages, 1)

		// capture capctity to compare later
		c := float64(cap(ages))
		// check for capactity condition
		if c != oldCap {
			fmt.Printf("len: %-10d capacity: %-10g, growth ratio: %.2f\n", len(ages), c, c/oldCap)
		}
		oldCap = c
	}
}
