package main

import "fmt"

// dividing into arbitrary sessions {session1}
func main() {
	// array of books
	var books [5]string
	books[0] = "Dracula"
	books[1] = "Adventures of Tom Sawyer"

	fmt.Printf("books: 			%#v\n", books)
	fmt.Printf("length:			%d\n", len(books))
	fmt.Printf("type: 			%T\n", books)
	fmt.Println()

	// slice of games
	// var games []string
	//  this would fail with the below error as the slice needs
	// to be initialized before using
	// panic: runtime error: index out of range [0] with length 0
	// games[0] = "Pokemon"
	// games[1] = "Full Metal Alchemy"

	// to counter the above, we can declare the slice as below
	games := []string{"Pokemon", "Full Metal Alchemy"}

	newGames := []string{"ping", "pong"}

	newGames = games

	games = nil
	// if slice is nil, go does not bother to run the range

	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not"
			break
		}
	}

	// if games = []string{} then it will not be nil
	games = []string{}
	// never check a slice is nil or not by comparing
	// value. alway go for length
	// if games == nil {
	// 	ok = "not"
	// }

	if len(games) != len(newGames) {
		ok = "not"
	}

	fmt.Printf("games: 			%#v\n", games)
	fmt.Printf("length: 		%d\n", len(games))
	fmt.Printf("type: 			%T\n", games)
	fmt.Printf("is slice nil? 	%t\n", games == nil)

	fmt.Printf("newGames: 		%#v\n", newGames)
	fmt.Printf("games and newGames are %s equal\n\n", ok)
}
