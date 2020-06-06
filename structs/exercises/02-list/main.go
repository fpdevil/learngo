package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item  item
	genre string
}

const commands = `
   > list	: Lists all the games
   > quit	: Quit

`

func main() {
	fmt.Println("---- 24-structs/02-list ----")
	fmt.Println("     EXERCISE: List         ")

	// use your solution from the previous exercise
	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{
			item:  item{id: 2, name: "x-com 2", price: 30},
			genre: "strategy",
		},
		{
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	fmt.Printf("Number of games in the game store: %d\n\n", len(games))

	fmt.Printf("%s\n", commands)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		cmd := in.Text()

		switch cmd {
		case "quit":
			fmt.Printf(">> %sting from the program\n", "quit")
			fmt.Println("   Bye!")
			return
		case "list":
			fmt.Printf(">> getting a %s of all games\n\n", "list")
			fmt.Printf("%-3s %-15s %-20s %-2s\n", "Id", "Name", "Genre", "Price")
			fmt.Printf("%s\n", strings.Repeat("-", 46))
			for _, v := range games {
				fmt.Printf("#%-2d %-15q %-20s $%-2d\n", v.item.id, v.item.name, "("+v.genre+")", v.item.price)
			}
			fmt.Printf("%s\n", commands)
		case "\n", "":
			fmt.Printf("%s\n", commands)
		}
	}

}
