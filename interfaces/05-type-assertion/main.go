package main

import "fmt"

func main() {
	fmt.Println("---- interfaces/05-type-assertion ----")
	fmt.Println()

	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubiks cube", price: 15}
	)

	// store all games + books into a store for availability
	var store list
	store = append(store, &minecraft, &tetris, &mobydick, &rubik)

	// display all the items
	store.print()

	// zero interface value is nil
	var p printer

	p = &tetris
	p.print()
}
