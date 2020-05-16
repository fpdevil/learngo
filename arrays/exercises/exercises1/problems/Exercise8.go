package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

// Exercise8 - Wizard Printer
func Exercise8() {
	fmt.Println()
	fmt.Println("---- 14-arrays/08-wizard-printer Exercise: Wizard Printer ----")

	// var scientists [...][3]string

	scientists := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%#v\n", scientists)
	sep := strings.Repeat("=", 50) + "\n"
	// fmt.Printf("%-17s %-17s %-14s", "First Name", "Last Name", "Nickname")
	for i, v := range scientists {
		if i == 1 {
			fmt.Printf("%s", sep)
		}
		fmt.Printf("%-17s %-17s %-14s\n", v[0], v[1], v[2])
	}
}
