package main

import "fmt"

type game struct {
	title string
	price money
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(rate float64) {
	g.price *= money(1 - rate)
}
