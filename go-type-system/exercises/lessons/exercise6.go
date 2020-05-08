package lessons

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert the Types
//
//  Convert the variables to appropriate types.
//
// EXPECTED OUTPUT
//  325.5 299.5
// ---------------------------------------------------------

//Exercise6 Convert the Types
func Exercise6() {
	fmt.Println()
	fmt.Println("---- 06-convert-the-types Exercise: Convert the Types ----")

	// DONT TOUCH THIS:
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	// DONT TOUCH THIS:
	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	// ----------------------------------------
	// FIX THE CODE BELOW:
	// You should solve it only by using conversions.
	// Do not change the code in any other way.

	// celsius *= celsiusDegree * factor
	// fahr *= fahrDegree * factor

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	// ----------------------------------------
	// DONT TOUCH THIS
	fmt.Println(celsius, fahr)

	// YOU MAY REMOVE THESE WHEN YOU'RE DONE
	_, _, _ = celsiusDegree, fahrDegree, factor
}
