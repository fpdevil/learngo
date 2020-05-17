package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// ---------------------------------------------------------
// EXERCISE: Ticker: Slide the Clock
//
//  Your goal is slide the placeholders every second.
//  Please run the solution to see it in action.
//
//
//  THIS IS A HARD EXERCISE:
//  + It will take you days but it will worth it.
//  + For experienced developers, this can take an hour or so.
//
//
//  1. You need to determine the starting and the ending digits to create
//     the sliding effect.
//
//
//  2. Each second, start from the next placeholder, skip the previous one.
//     This means: Only draw the next placeholders.
//
//     Like this:
//
//     12:40:31
//     2:40:31
//     40:31
//     0:31
//     :31
//     31
//     1
//
//
//  3. After the last placeholder is displayed, fill the lines for the missing
//     placeholders, and then start from the first placeholder. Draw it to the
//     right part of the screen.
//
//     Like this:
//
//     12:40:31
//     2:40:31
//     40:31
//     0:31
//     :31
//     31
//     1
//            1
//           12
//          12:
//         12:4
//        12:40
//       12:40:
//      12:40:3
//     12:40:31
//
//     As you can see, you need to draw the clock from the right part of the
//     screen, beginning from the first placeholder.
//
//
// HINTS
//   + You would need to clear the screen inside the loop instead of once.
//     Otherwise the previous placeholders will be left on the screen.
//
//
// EXPECTED OUTPUT
//   Please run the solution to see it in action. Do not look at the source-code
//   though.
// ---------------------------------------------------------

// go run *.go

func main() {
	// clear the screen first
	screen.Clear()

	// infinite loop
	for {
		screen.MoveTopLeft()

		t := time.Now()
		hour, minute, second := t.Hour(), t.Minute(), t.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			sep,
			digits[minute/10], digits[minute%10],
			sep,
			digits[second/10], digits[second%10],
		}

		// for row := range clock[0] {
		// 	for index, digit := range clock {
		// 		// handle blinking of separator
		// 		next := clock[index][row]
		// 		if digit == sep && second%2 == 0 {
		// 			next = "   "
		// 		}
		// 		fmt.Printf("%6s", next)
		// 	}
		// 	fmt.Println()
		// }

		for row := range clock[0] {
			for index, digit := range clock {
				// handle blinking of separator
				next := clock[index][row]
				if digit == sep && second%2 == 0 {
					next = "   "
				}
				fmt.Printf("%6s", next)
			}
			fmt.Println()
		}

		time.Sleep(time.Second / 10)
	}
}
