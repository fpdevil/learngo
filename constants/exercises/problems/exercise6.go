package problems

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: No Conversions Allowed
//
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
//
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------

// Exercise6 handles No Conversions Allowed
func Exercise6() {
	fmt.Println("")
	fmt.Println("---- 10-constants/06-no-conversions-allowed Exercise: No Conversions Allowed ----")

	// const later int = 10
	// hours, _ := time.ParseDuration("1h")
	// fmt.Printf("%s later...\n", hours*later)

	const later int = 10
	hours, _ := time.ParseDuration("1h")
	fmt.Printf("%s later...\n", hours*time.Duration(later))
}
