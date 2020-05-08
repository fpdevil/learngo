package main

import (
	"fmt"

	"github.com/fpdevil/learngo/go-type-system/underlying-types/weights"
)

// Name type declaration
type (
	// Gram whose underlying type is int
	Gram int

	// Kilogram whose udnerlying type is int
	Kilogram int

	// Ton whose underlying type is int
	Ton int
)

func main() {
	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	// types with different names cannot be used together,
	// eventhough their underlying types are same
	// salt = apples => illegal

	// but, since their underlying type is still the same "int",
	// they can be converted to any defined type with an underlying
	// type of "int"
	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Kilogram(Gram(apples)))

	fmt.Printf("salt: %d [%[1]T], apples: %d [%[2]T], trucks: %d [%[3]T]\n", salt, apples, truck)

	//  TYPES FROM OTHER PACKAGES
	// gram and weigths.Gram are from different packages
	// gram = weights.Gram
	salt = Gram(weights.Gram(100))

	fmt.Printf("Type of weights.Gram is %T\n", weights.Gram(100))
	fmt.Printf("Type of main.Gram is %T\n", Gram(100))

	// we may declare a new type
	// from a type of an external package
	type myGram weights.Gram

	var pepper myGram = 99
	pepper = myGram(salt)

	fmt.Printf("Type of pepper is %T\n", pepper)
}
