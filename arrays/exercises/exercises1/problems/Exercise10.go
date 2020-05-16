package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

// Exercise10 - Hipster's Love Bookstore Search Engine
func Exercise10(vals string) {
	fmt.Println()
	fmt.Println("---- 14-arrays/10-hipsters-love-search Exercise: Hipster's Love Bookstore Search Engine ----")

	args := strings.Split(vals, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if vals == "" || len(args) < 1 {
		fmt.Println("Tell me a book title")
		return
	}

	query := args[0]
	var booklist []int

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	for i, v := range books {
		book := strings.ToLower(v)
		squery := strings.ToLower(query)

		if strings.Contains(book, squery) {
			booklist = append(booklist, i)
		}
	}

	fmt.Println("Search Results:")
	if len(booklist) != 0 {
		for _, v := range booklist {
			fmt.Printf("+ %s\n", books[v])
		}
		return
	}

	fmt.Printf("We don't have the book: %q\n", query)
}
