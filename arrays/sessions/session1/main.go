package main

import "fmt"

func main() {
	var books [4]string
	fmt.Printf("books type: %T\n", books) // books type: [4]string
	fmt.Printf("books: %q\n", books)      // books: ["" "" "" ""]
	fmt.Printf("books: %#v\n", books)     // books: [4]string{"", "", "", ""}

	// comparing the arrays - types should be identical
	// each array should contain same number of elements
	var (
		blue  = [3]int{6, 9, 3}
		red   = [3]int{6, 9, 3}
		green = [...]int{3, 6, 9}
	)

	fmt.Printf("blue bookcase: %#v\n", blue)
	fmt.Printf("red bookcase: %#v\n", red)
	fmt.Printf("green bookcase: %#v\n", green)
	fmt.Println("comparing red & blue:", blue == red)
	fmt.Println("comparing red & green:", red == green)

	// assigning one array to another
	prev := [3]string{
		"Kafkas Revenge",
		"Stay Golden",
		"Everythingship",
	}
	fmt.Printf("last year collection: %#v\n", prev)

	var current [4]string
	for i, b := range prev {
		current[i] += b + " 2nd Edition."
	}
	current[3] = "New Revolution"
	fmt.Printf("this year collection: %#v\n", current)
}
