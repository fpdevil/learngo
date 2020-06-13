package main

import "fmt"

type game struct {
	title string
	price float64
}

// #1a
// func (g game) print() {
// 	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
// }

// change #1a also to a pointer type forconsitency
// #1b
func (g *game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

// offer discount for A GAME

// #2a
// g here is a copy, changing the value of g here,
// does not change the price of the original game variable
// func (g game) discount(ratio float64) {
// 	g.price *= (1 - ratio)
// 	// debugging
// 	g.print()
// 	fmt.Println("debugging discount()")
// }

// 2b
// chaning g to a pointer type
func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
	// debugging
	// g.print()
	// fmt.Println("debugging discount()")
}
