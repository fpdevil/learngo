package main

import "fmt"

func main() {
	fmt.Println("--- 24-structs/04-embedding ---")
	
	type text struct {
		title string
		words int
	}

	// type book struct {
	// 	title string
	// 	words int
	// 	isbn string

	// 	// including text as field
	// 	text text
	// }

	// moby := book{
	// 	title: "moby dick",
	// 	words: 206052,
	// 	isbn: "102030",
	// }
	// fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

	// type book struct {
	// 	// title string
	// 	// words int
	// 	isbn string

	// 	// including text as field
	// 	text text
	// }

	// tintin := book{
	// 	text: text{title: "adventures of tintin", words: 161250},
	// 	isbn: "581276",
	// }

	// fmt.Printf("%s has %d words (isbn: %s)\n", tintin.text.title, tintin.text.words, tintin.isbn)

	type book struct {
		// title string
		// words int

		// including text as field
		// text text

		// embed as text
		text

		isbn string

		// now insert the conflicting field
		title string
	}

	ts := book {
		text: text{ title: "adventures of tom sawyer", words: 21007},
		isbn: "812756",
	}
	fmt.Printf("%s has %d words (isbn: %s)\n", ts.text.title, ts.text.words, ts.isbn)

	ts.text.words = 999
	ts.words++
	fmt.Println("after the text.words update...")
	fmt.Printf("%s has %d words (isbn: %s)\n", ts.text.title, ts.words, ts.isbn)

	moby := book{
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	// title will not work here
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)
	// change to moby.text.title
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.text.title, moby.words, moby.isbn)
	fmt.Printf("%#v\n", moby)
}
