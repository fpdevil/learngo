// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	fmt.Println("---- 26-pointers/01-basics ----")
	fmt.Println("     EXERCISE: Basics			")

	// create a nil pointer with the type of pointer to a computer
	var P *computer

	// compare the pointer variable to a nil value, and say it's nil
	if P == nil {
		fmt.Printf("computer %v is nil and it's address is: %[1]p\n", P)
	}

	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}

	// put the apple into a new pointer variable
	newApple := apple

	// compare the apples: if they are equal say so and print their addresses
	if apple == newApple {
		fmt.Printf("Apples are equal with addresses          : apple: %p & new apple: %p\n", apple, newApple)
	}

	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if apple != sony {
		fmt.Printf("apple and sony are not equal             : apple: %p & sony: %p\n", apple, sony)
	}

	// put apple's value into a new ordinary variable
	appleSys := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple                                    : %p, %p\n", &apple, apple)
	fmt.Printf("appleSys                                 : %p\n", &appleSys)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple == appleSys {
		fmt.Println("apple and appleSys are equal")
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("apple                                    : %p - appleVal: %p\n", &apple, apple)
	fmt.Printf("apple                                    : %+v - appleVal: %+v\n", *apple, appleSys)

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Printf("sony                                     : %s\n", sony.brand)

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Printf("returned value of apple                  : %+v\n", retValue(apple))
	fmt.Printf("returned value of sony               	 : %+v\n", retValue(sony))

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("apple                                    : %p\n", newConstructor("apple"))
	fmt.Printf("lenovo                                   : %p\n", newConstructor("lenovo"))
	fmt.Printf("dell                                     : %p\n", newConstructor("dell"))
}

// create a new function: change
// the func can change the given computer's brand to another brand
func change(c *computer, brand string) {
	c.brand = brand
}

// write a func that returns the value that is pointed by the given *computer
// print the returned value
func retValue(c *computer) computer {
	return *c
}

func newConstructor(brand string) *computer {
	return &computer{brand: brand}
}
