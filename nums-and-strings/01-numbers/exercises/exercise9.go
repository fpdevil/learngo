package exercises

import (
	"fmt"
	"math"
)

// ---------------------------------------------------------
// EXERCISE: Sphere Volume
//
//  1. Get the radius from the command-line
//  2. Convert it to a float64
//  3. Calculate the volume of a sphere
//
// SPHERE VOLUME FORMULA
//  https://en.wikipedia.org/wiki/Sphere#Enclosed_volume
//
// RESTRICTION
//  Use `math.Pow` to calculate the volume
//
// EXPECTED OUTPUT
//  go run main.go 10
//    4188.79
//
//  go run main.go .5
//    0.52
// ---------------------------------------------------------

// Exercise9 from the 09-sphere-volume
func Exercise9(r float64) {
	fmt.Println()
	fmt.Println("--- EXERCISE: Sphere Volume ---")

	var radius, vol float64

	// ADD YOUR CODE HERE
	// ...
	radius = r
	vol = (4 * math.Pi * math.Pow(radius, 3)) / 3

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)
}
