package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("---- interfaces/02-receivers ----")
	fmt.Println()

	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// calling methods using the types is also possible
	book.print(mobydick)
	fmt.Printf("%s\n", strings.Repeat("-", 10))

	minecraft.discount(0.5) // same as (&minecraft).discount(0.5)
	fmt.Printf("%s\n", strings.Repeat("-", 10))

	mobydick.print()
	minecraft.print()
	tetris.print()

	// creates a variable that occupies 200mb on memory
	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}
