package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

// Exercise3 - Refactor to Array Literals
func Exercise3() {
	fmt.Println("")
	fmt.Println("---- 14-arrays/03-array-literal Exercise: Refactor to Array Literals ----")

	// args := strings.Split(magstr, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// if magstr == "" {
	// 	fmt.Println("Give me the magnitude of the earthquake.")
	// 	return
	// }

	names := [3]string{
		"Mickey Mouse",
		"Donald Duck",
		"Asterix",
	}

	distances := [5]int{
		10, 20, 30, 40, 50,
	}

	data := [5]byte{'H', 'O', 'W', 'D', 'Y'}

	ratio := [1]float64{1.71324}

	alives := [5]bool{
		true, true, false, true, false,
	}

	var zero [0]byte // this doesn't work zero := [0]byte{}

	sep := "\n" + strings.Repeat("=", 20) + "\n"
	fmt.Print("\nnames", sep)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", sep)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", sep)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", sep)
	for i := 0; i < len(ratio); i++ {
		fmt.Printf("ratio[%d]: %.2f\n", i, ratio[i])
	}

	fmt.Print("\nalives", sep)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Print("\nzero", sep)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}

	fmt.Printf(`
%s
FOR RANGES
%[1]s
	`, strings.Repeat("~", 30))

	fmt.Print("\nnames", sep)
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}

	fmt.Print("\ndistances", sep)
	for i, v := range distances {
		fmt.Printf("distances[%d]: %d\n", i, v)
	}

	fmt.Print("\ndata", sep)
	for i, v := range data {
		fmt.Printf("data[%d]: %d\n", i, v)
	}

	fmt.Print("\nratios", sep)
	for i, v := range ratio {
		fmt.Printf("ratio[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nalives", sep)
	for i, v := range alives {
		fmt.Printf("alives[%d]: %t\n", i, v)
	}

	fmt.Print("\nzero", sep)
	for i, v := range zero {
		fmt.Printf("zero[%d]: %d\n", i, v)
	}
}
