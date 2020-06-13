package main

import "fmt"

func main() {
	fmt.Println("---- interfaces/01-methods ----")
	fmt.Println()

	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	// #1
	// printBook(mobydick)
	// printGame(minecraft)
	// printGame(tetris)

	// #2
	// mobydick.printBook()
	// minecraft.printGame()
	// tetris.printGame()

	// #3
	mobydick.print()
	minecraft.print()
	tetris.print()
}
