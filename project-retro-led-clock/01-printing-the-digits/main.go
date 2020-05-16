package main

import "fmt"

// 1. Create each digit as an array
// 2. Clear the screen
// 3. Create an infinite loop
// 		- Move cursor to the top left
// 		- Update the clock array by copying the digits to it
// 		- Draw clock using clock array
// 		- Sleep for a second

func main() {
	fmt.Printf("GOAL 1: Printing the Digits...\n")
	// [x] Goal 1 steps
	// [x] Define a new placeholder type
	// [x] Create the digit arrays from 0 to 9
	// [x] Put them into the "digits" array
	// [x] Print them side-by-side

	// defining a new placeholder type
	type placeholder = [5]string

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
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	fmt.Printf("Place holder type for digits: %T\n", digits)
	fmt.Println("")

	// // first element
	fmt.Println("Priting only first digit...")
	for _, row := range digits[0] {
		fmt.Println(row)
	}
	fmt.Println()

	fmt.Println("Printing digits vertically...")
	for _, row := range digits {
		for _, v := range row {
			fmt.Println(v)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Using field length...")
	for row := range digits[0] {
		for digit := range digits {
			fmt.Printf("%6s", digits[digit][row])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Using a print format...")
	// loop each of the 5 rows of each digit
	for row := range digits[0] {
		// loop over each of the 10 strings in the digits
		for digit := range digits {
			// after each first element of digit place a space
			fmt.Print(digits[digit][row], "  ")
		}
		// place a new line after each row
		fmt.Println()
	}
}
