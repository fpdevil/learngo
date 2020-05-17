package main

import (
	"fmt"
	"time"

	// import this package for screen clearing
	"github.com/inancgumus/screen"
)

// ---------------------------------------------------------
// EXERCISE: Add Split Seconds
//
//  Your goal is adding the split second to the clock. A split second is
//  1/10th of a second.
//
//  1. Find the current split second
//  2. Add dot character to the clock (as in the expected output)
//  3. Add the split second digit to the clock
//  4. Blink the dot every two seconds (just like the separators)
//  5. Update the clock every 1/10th of a second, instead of every second.
//     (Update the clock every 100 millliseconds)
//
// HINTS
//   + You can find the split second using Nanosecond method of the Time type.
//     https://golang.org/pkg/time/#Time.Nanosecond
//
//   + A split second is the first digit of the Nanosecond.
//
//   + Remember: time.Second is an integer constant, so it can be divided
//               with a number.
//
//     https://golang.org/pkg/time/#Time.Second
//
// EXPECTED OUTPUT
//     Note that, clock is updated every split second instead of a second.
//
//     Separators are displayed (second is an odd number):
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       █ █
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ██
//      █    █    ░   █     █    ░    █     █        █
//      █    █        ███   █         █     █        █
//      █    █    ░     █   █    ░    █     █        █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █         █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █         █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       █ █
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░     █
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █         █
//      █    █        ███   █         █     █         █
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░     █
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █       █ █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █    ░   █     █    ░    █     █       █ █
//      █    █        ███   █         █     █       ███
//      █    █    ░     █   █    ░    █     █         █
//     ███  ███       ███  ███       ███    █   ░   ███
//
//     Separators are not displayed (second is an even number):
//
//     ██   ██        ███  ██        ██   ███       ███
//      █    █        █     █         █   █ █       █ █
//      █    █        ███   █         █   ███       █ █
//      █    █          █   █         █   █ █       █ █
//     ███  ███       ███  ███       ███  ███       ███
//
// ---------------------------------------------------------

// go run *.go

func main() {
	// clear the screen first for once
	screen.Clear()

	// infinite loop
	for {
		// move the cursor to the top left
		screen.MoveTopLeft()

		t := time.Now()
		hour, minute, second := t.Hour(), t.Minute(), t.Second()
		splitsecond := t.Nanosecond() / 1e08

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			sep,
			digits[minute/10], digits[minute%10],
			sep,
			digits[second/10], digits[second%10],
			period,
			digits[splitsecond],
		}

		for row := range clock[0] {
			for index, digit := range clock {
				// handle the blinking of separator
				next := clock[index][row]
				if (digit == sep || digit == period) && second%2 == 0 {
					next = "   "
				}
				fmt.Printf("%6s", next)
			}
			fmt.Println()
		}
		time.Sleep(time.Second / 10)
	}
}
