package main

import "fmt"

func main() {
	fmt.Println("---- interfaces/03-nonstructs ----")
	fmt.Println()

	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	mobydick.print()
	minecraft.print()
	tetris.print()

	var items []*game
	items = append(items, &minecraft, &tetris)

	// we can attach methods to a compatible type on the fly:
	// items -> []game
	// items -> []*game
	my := list(items)
	my = nil

	// you can call methods even on a nil value
	my.print()

}
