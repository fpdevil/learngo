package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id int, name string, price int, genre string) game {
	item := item{
		id:    id,
		name:  name,
		price: price,
	}
	g := game{
		item:  item,
		genre: genre,
	}
	return g
}

func addGame(games []game, newgame game) []game {
	games = append(games, newgame)
	return games
}

func load() []game {
	// var games []game
	game1 := newGame(1, "god of war", 50, "action adventure")
	game2 := newGame(2, "x-com 2", 40, "strategy")
	game3 := newGame(3, "minecraft", 20, "sandbox")

	games := addGame([]game(nil), game1)
	games = addGame(games, game2)
	games = addGame(games, game3)

	return games
}

func indexByID(games []game) map[int]game {
	indexByID := make(map[int]game)
	for _, g := range games {
		indexByID[g.id] = g
	}
	return indexByID
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
}
