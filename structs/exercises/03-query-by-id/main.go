package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
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
   > id N	: Queries a game by id
   > quit	: Quit

`

func main() {
	fmt.Println("---- 24-structs/03-query-by-id ----")
	fmt.Println("     EXERCISE: Query By Id  		")

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

	var gameByID map[int]game
	gameByID = make(map[int]game)

	for _, g := range games {
		gameByID[g.item.id] = g
	}

	fmt.Printf("Number of games in the game store: %d\n\n", len(games))

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s\n", commands)

		// if the user does not type anything, continue
		if !in.Scan() {
			break
		}

		cmd := strings.Fields(in.Text())
		lcd := len(cmd)

		if lcd == 0 {
			continue
		}

		switch cmd[0] {
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
		case "id":
			if lcd != 2 {
				fmt.Print("wrong id\n")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := gameByID[id]

			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}
			fmt.Printf("#%-2d %-15q %-20s $%-2d\n", g.item.id, g.item.name, "("+g.genre+")", g.item.price)

		}
	}
}
