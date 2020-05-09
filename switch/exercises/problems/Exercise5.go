package problems

import (
	"fmt"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

// Exercise5 - Days in a Month
func Exercise5(mstr string) {
	fmt.Println("")
	fmt.Println("---- 12-switch/05-days-in-month Exercise: Days in a Month ----")

	args := strings.Split(mstr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if mstr == "" && len(mstr) != 1 {
		fmt.Println("Give me a month name.")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, args[0]

	switch m := strings.ToLower(month); m {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
