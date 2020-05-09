package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month!")
		return
	}

	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter Season!")
		// break is not needed as it's automatically
		// included by go
	case "Mar", "Apr", "May":
		fmt.Println("Spring Season!")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer Season!")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall Season!")
	// the order of default does not matter and it
	// can occur or be placed anywhere (even first)
	default:
		fmt.Println("%q is not a month.\n", m)
	}

	// fallthrough
	i := 142

	// fallthrough statement should be the last statement
	// inside a case clause
	// fallthrough is a sort of replacement for break

	switch {
	case i > 100:
		// fmt.Println("big positive number")

		// using fallthrough statement to tackle
		// deuplicate repeated number
		fmt.Println("big ")
		fallthrough
	case i > 0:
		// fmt.Println("Postive number")
		fallthrough
	default:
		fmt.Println("number")
	}

	// parts of a day
	// t := time.Now()
	// hour := t.Hour()
	// fmt.Println("current hour is %d", hour)

	switch hour := time.Now().Hour(); {
	case hour >= 18:
		fmt.Println("Good Evening!")
	case hour >= 12:
		fmt.Println("Good Afternoon!")
	case hour >= 6:
		fmt.Println("Good Morning!")
	default:
		fmt.Println("Good Night!")
	}
}
