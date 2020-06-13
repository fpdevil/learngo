package main

import "fmt"

/*
	We cannot have multiple functions or multiple methods with
  same name within the same package
*/

type game struct {
	title string
	price float64
}

// #1
// here is g is a copy of the game
// func printGame(g ga) {
// 	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
// }

// #2
// func (g game) printGame() {
// 	// here g is still a copy of the original `game` value
// 	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
// }

// #3
func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
