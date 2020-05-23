package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("pagination ...")

	// think of this as search results of a search engine.
	// it could have been fetched from a database
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	l := len(items)
	const pageSize = 4

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		fmt.Printf("[%d:%d] / %d\n", from, to, l)
		if to > l {
			to = l
		}

		currentPage := items[from:to]

		// put a header for each value
		header := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		s.Show(header, currentPage)
	}
}
