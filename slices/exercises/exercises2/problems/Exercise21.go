package problems

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

// Exercise21 - Correct the lyric
func Exercise21() {
	fmt.Println("")
	fmt.Println("---- 16-slices/21-correct-the-lyric Exercise: Correct the lyric ----")

	// args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

	// ADD YOUR CODE BELOW:
	// ...
	lyric = append([]string{"yesterday"}, lyric...)

	s.PrintBacking = true
	// s.Show("lyric", lyric)

	// s.Show("lyric[:8]", lyric[:8])
	// s.Show("lyric[8+5:]", lyric[8+5:])
	// s.Show("lyric[8:13]", lyric[8:13])

	// the portion of the lyric "oh i believe in yesterday" spans
	// from 8 till 13th index, so let's grab that and first append
	// it to the end there by duplicating the values.
	//--- this created a new backing array ---
	lyric = append(lyric, lyric[8:13]...)
	// s.Show("Duplicate 8 to 13", lyric)

	// next, since the same lyric is existing from 8 to 13, we can
	// override that with the lyric string from 13 till end
	//--- this used the previous new backing array ---
	lyric = append(lyric[:8], lyric[13:]...)
	// s.Show("Append till 8 From 13 to end", lyric)

	fmt.Printf("%s\n", lyric)
}
