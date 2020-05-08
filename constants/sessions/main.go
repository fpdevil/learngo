package main

import "fmt"

func main() {
	// centimeters into meters
	const meters int = 100 // named constant to replace magic number
	// This program uses magic values
	cm := 100
	// 100 is a magic value, its a value without a name
	// m := cm / 100
	m := cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	// 100 is a magic value, its a value without a name
	// m = cm / 100
	m = cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	fmt.Println("Using IOTA...")
	const (
		EST = -(5 + iota) // results in -(5 + 0) = -5
		_                 // blank identifier to keep iota count
		MST               // keeps count at -7
		PST               // keeps count at -8
	)
	fmt.Printf("EST: %d (%[1]T), MST: %d (%[2]T), PST: %d (%[3]T)", EST, MST, PST)
}
