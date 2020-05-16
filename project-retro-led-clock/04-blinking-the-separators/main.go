package main

import (
	"fmt"
	"time"

	// package for clearing the screen
	"github.com/inancgumus/screen"
)

// GOAL 4: Blinking the Separators

func main() {
	fmt.Println("GOAL 4: Blinking the Separators")

	// named type for holding the digits
	type placeholder = [5]string

	// Define  the  digits from  0  through  9 as  strings  of  size 5  using
	// the  character █.  The  placeholder  can be  used  to represent  the
	// complete digit  set using a multidimensional  array of size 10  X 5 as
	// [10][5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}

	// const sep = "░"
	sep := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// clear all the characters on the screen
	screen.Clear()

	// creting an infinite loop for animating
	for {
		// Moves the cursor to the top-left position of the screen
		screen.MoveTopLeft()

		// get the current time
		t := time.Now()
		hour, minute, second := t.Hour(), t.Minute(), t.Second()

		// fmt.Printf("Time: 	%v:%v:%v\n", hour, minute, second)

		// and seconds for  clock time with separator, we  need eight placeholder
		// create a clock array using  placeholder matching the time hour, minute
		// values for 6 digits and 2 separators
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			sep,
			digits[minute/10], digits[minute%10],
			sep,
			digits[second/10], digits[second%10],
		}

		for row := range clock[0] {
			for digit := range clock {
				// move the time string to a variable first
				update := clock[digit][row]
				// update the time string to include an empty fields
				// for every 2 seconds if the specific digit is a separator
				if clock[digit] == sep && second%2 == 0 {
					update = "   "
				}
				fmt.Printf("%5s", update)
			}
			fmt.Println()
		}
		fmt.Println()
		// sleep for every 1 second
		time.Sleep(time.Second * 1)
	}
}
