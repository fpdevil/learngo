package main

import (
	"fmt"
	"time"
)

// STEP #2 — Print the Clock
//  Get the current time
//  Create the clock array
//  Print the clock
//  Add the separators

func main() {
	fmt.Println("GOAL 2: Printing the Clock")
	// name type fot holding the digits
	type placeholder = [5]string

	// Define the digits from 0 through 9 as strings
	// of size 5. The placeholder can be used to represent
	// the complete digit set using a multidimensional
	// array of size 10 X 5 as [10][5]string
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

	// get the current time
	t := time.Now()
	hour, minute, second := t.Hour(), t.Minute(), t.Second()

	fmt.Printf("Time: 	%v:%v:%v\n", hour, minute, second)
	fmt.Println()

	// create a clock array using placeholder matching the
	// time hour, minute and seconds
	// for clock time with separator, we need eight placeholder
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
			fmt.Printf("%5s", clock[digit][row])
		}
		fmt.Println()
	}
}
