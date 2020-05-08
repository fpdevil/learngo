package lessons

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Optimal Types
//
//  1. Choose the optimal data types for the given situations.
//  2. Print them all
//  3. Try converting them to lesser data types.
//     For example try converting int64 variable to int32.
//     Then observe the result.
//     Search the web why the result is so?
//
// NOTE
//  This is just an exercise for teaching you different
//  data types. Do not apply it to the real-life.
//
//  As I said in the lectures that, premature optimization
//  is not a good thing.
// ---------------------------------------------------------

//Exercise1 Optimal Types
func Exercise1() {
	fmt.Println()
	fmt.Println("---- 01-optimal-types Exercise: Optimal Types ----")

	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	var letter byte = 'S'
	fmt.Println("an english letter (byte): ", letter)

	// a non-english letter (search web for: unicode codepoint)
	var unicode = 'λ' // var unicode byte = 'λ'
	fmt.Println("a non-english letter (rune): ", unicode)

	// a year in 4 digits like 2040
	var year uint16 = 2040
	fmt.Println("a year in 4 digits like 2040 (uint16): ", year)

	// a month in 2 digits: 1 to 12
	var month uint8 = 9
	fmt.Println("a month in 2 digits: 1 to 12 (uint8): ", month)

	// the speed of the light
	var speed uint32 = 299792458
	fmt.Println("the speed of the light (uint32): ", speed, "m/sec")

	// angle of a circle
	var angle float32 = 42.5
	fmt.Println("angle of a circle (float32):", angle, "degrees")

	// PI number: 3.141592653589793
	// var pi float64 = 3.141592653589793
	var pi = 3.141592653589793
	fmt.Println("PI number (float64): ", pi)
}
