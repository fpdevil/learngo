package problems

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Add a newline after each sentence
//
//  You have a slice that contains Beatles' awesome song:
//  Yesterday. You want to add newlines after each sentence.
//
//  So, create a new slice and copy every words into it. Lastly,
//  after each sentence, add a newline character ('\n').
//
//
// ORIGINAL SLICE:
//
//   [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
// EXPECTED SLICE (NEW):
//
//   [yesterday all my troubles seemed so far \n away now it looks as though they are here to stay \n oh i believe in yesterday \n]
//
//
// CURRENT OUTPUT
//
//  yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday
//
// EXPECTED OUTPUT
//
//  yesterday all my troubles seemed so far away
//  now it looks as though they are here to stay
//  oh i believe in yesterday
//
//
// RESTRICTIONS
//
//  + Don't use `append()`, use `copy()` instead.
//
//  + Don't cheat like this:
//
//    fmt.Println(lyric[:8])
//    fmt.Println(lyric[8:18])
//    fmt.Println(lyric[18:23])
//
//  + Create a new slice that contains the sentences
//    with line endings.
//
//
// NOTE
//
// If the program does not work, please update your
// local copy of the prettyslice package:
//
//   go get -u github.com/inancgumus/prettyslice
//
//
// ---------------------------------------------------------

// Exercise25 - Add a newline after each sentence
func Exercise25() {
	fmt.Println("")
	fmt.Println("---- 16-slices/25-add-lines Exercise: Add a newline after each sentence ----")

	// args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// You need to add a newline after each sentence in another slice.
	// Don't touch the following code.
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday`)

	// ===================================
	//
	// ~~~ CHANGE THIS CODE ~~~
	//
	// fix := lyric
	//
	// define a slice of length = len(lyric) + 3 newline chars
	fix := make([]string, len(lyric)+3)

	// define a list with stopping points where newline can be inserted
	indexl := []int{8, 18, 23}

	// define variable to hold previous value
	var previous int

	for i, current := range indexl {

		// fmt.Printf("previous: %d, current: %d\n", previous, current)
		// fmt.Printf("lyric[%d:%d] = %s\n", previous, current, lyric[previous:current])
		// fmt.Printf("%d - newline position: %d\n", i, current+i)
		// adjust the offset value equal to iteration value for
		// each iteration which is the value of i
		copy(fix[previous+i:current+i], lyric[previous:current])

		// as per Inanc comment (updated the copy of \n)
		// You don't have to copy there, just change the slice's
		// element by using an index expression.
		// copy(fix[current+i:], []string{"\n"})
		fix[current+i] = "\n"

		// make the previous value same as current
		previous = current
	}
	//
	// ===================================

	// Currently, it prints every sentence on the same line.
	// Don't touch the following code.
	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
