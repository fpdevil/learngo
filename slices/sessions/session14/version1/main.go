package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("Multidimensional Slices...")

	// spendings on each day
	// 1st day: $200, $100
	// 2nd day: $500
	// 3rd day: $50, $25, and $75
	spendings := [][]int{
		{200, 100},
		{500},
		{50, 25, 75},
	}

	for i, daily := range spendings {
		var total int
		for _, spent := range daily {
			total += spent
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}
