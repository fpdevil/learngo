package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

const commands = `
   > list	: Lists all the games
   > id N	: Queries a game by id
   > save 	: Encode the games to json
   > quit	: Quit

`

func main() {
	fmt.Println("---- 24-structs/05-decode ----")
	fmt.Println("     EXERCISE: Decode		   ")

	// use your solution from the previous exercise
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item  item
		genre string
	}

	// create a new exportable json struct for Encoding/Decoding
	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	//  1. Load the initial data to the game store from json.
	var jsonData []jsonGame
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	//  2. Load the decoded values into the usual `game` values
	//     (to the games slice as well).
	var games []game
	for _, jd := range jsonData {
		games = append(games, game{item{jd.ID, jd.Name, jd.Price}, jd.Genre})
	}

	// ======================================================
	// games := []game{
	// 	{
	// 		item:  item{id: 1, name: "god of war", price: 50},
	// 		genre: "action adventure",
	// 	},
	// 	{
	// 		item:  item{id: 2, name: "x-com 2", price: 30},
	// 		genre: "strategy",
	// 	},
	// 	{
	// 		item:  item{id: 3, name: "minecraft", price: 20},
	// 		genre: "sandbox",
	// 	},
	// }
	// ======================================================

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
		case "save":
			// now create a new slice for the new struct
			jsonGames := []jsonGame{}

			// populate the data from games into jsonGames
			for _, g := range games {
				jsonGames = append(jsonGames, jsonGame{g.item.id, g.item.name, g.genre, g.item.price})
			}

			bytes, err := json.MarshalIndent(jsonGames, "", "\t")
			if err != nil {
				fmt.Println("invalid data")
				continue
			}

			fmt.Printf("%s\n", string(bytes))
			return
		}
	}
}
