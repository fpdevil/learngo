package main

import "fmt"

func main() {
	fmt.Println("---- interfaces/04-interfaces ----")
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

	// interface values can be comparable
	fmt.Println(store[0] == &minecraft)
	fmt.Println(store[3] == &rubik)
}
