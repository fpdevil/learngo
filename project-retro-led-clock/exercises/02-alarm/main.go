package main

import (
	"fmt"
	"time"

	// import this package for screen clear
	"github.com/inancgumus/screen"
)

// ---------------------------------------------------------
// EXERCISE: Set an Alarm
//
//  Goal is printing " ALARM! " every 10 seconds.
//
// EXPECTED OUTPUT
//
//     ██   ███       ███  ██        ███  ███
//      █     █   ░     █   █    ░     █  █ █
//      █   ███       ███   █        ███  ███
//      █   █     ░     █   █    ░     █    █
//     ███  ███       ███  ███       ███  ███
//
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  ███   █
//          ███  █    ███  ██   █ █   █
//          █ █  █    █ █  █ █  █ █
//          █ █  ███  █ █  █ █  █ █   █
//
//     ██   ███       ███  ██        █ █  ██
//      █     █   ░     █   █    ░   █ █   █
//      █   ███       ███   █        ███   █
//      █   █     ░     █   █    ░     █   █
//     ███  ███       ███  ███         █  ███
//
// HINTS
//
//  <<< BEWARE THE SPOILER! >>>
//
//  I recommend you to first solve the exercise yourself before reading the
//  following hint.
//
//
//  + You can create a new array named alarm with the same length of the
//    clocks array, so you can copy the alarm array to the clocks array
//    every 10 seconds.
//
// ---------------------------------------------------------

// go run *.go

func main() {
	// clear the screen once.
	screen.Clear()

	// infinite loop
	for {
		// move the cursor to the top left
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

		// trigger alarm at every 10 seconds (we will get alarm for every 10th sec)
		callalarm := second%10 == 0

		if callalarm {
			clock = alarm
		}

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
		time.Sleep(time.Second)
	}
}
