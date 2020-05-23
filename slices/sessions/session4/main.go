package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("slicing...")

	// think of this as search results of a search engine.
	// it could have been fetched from a database
	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	// show only 4 items per line
	s.MaxPerLine = 4
	s.Show("All the items", items)

	top3 := items[:3]
	s.Show("Top 3 items", top3)

	l := len(items)
	// we can use variables within slice expressions as under
	last4 := items[l-4:]
	s.Show("Last 4 items", last4)

	// slicing already sliced element
	// this is reslicing
	middle := last4[1:3]
	s.Show("Las4[1;3]", middle)

	// it's possible to have same elements in different indexes
	s.Show(items[9], last4[0])
}
