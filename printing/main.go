package main

import "fmt"

func main() {
	var brandx string

	// prints the string in quoted-form like this ""
	fmt.Printf("%q\n", brandx)

	brandx = "Tesla"
	fmt.Printf("%q\n", brandx)

	ops, ok, fail := 2350, 543, 433
	fmt.Println("total:", ops, "- success:", ok, "/", fail)
	// using printf
	fmt.Printf("total: %d - success: %d / %d\n", ops, ok, fail)

	// escape sequences (\n escape sequence newline)
	fmt.Println("hi\nhi")
	fmt.Printf("hi\n\"hi\"")
	fmt.Println("--- some more examples for formatting ---")

	var (
		speed int
		heat  float64
		off   bool
		brand string
	)
	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	// %v verb is the swiss army knife
	// it can print any type of value but not type safe
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does: %v has life? %v\n", planet, hasLife)

	// argument indexing - indexes start from 1
	fmt.Println("--- with argument indexing ---")
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG!\n",
		planet,
		distance,
	)

	fmt.Println()
	fmt.Println("--- more examples with type safety ---")
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Does: %s has life? %t\n", planet, hasLife)

	fmt.Println()
	fmt.Printf("--- \"orbital with precision\" ---\n")
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
	fmt.Printf("Orbital Period: %.5f days\n", orbital)

	fmt.Println()
	fmt.Printf("--- \"general examples\" ---\n")
	fmt.Printf("The value %.2f is %[1]T and value %.3f is %[2]T\n", 3.414, 1.732)

	fmt.Printf("--- \"print binary numbers\" ---\n")
	for i := 0; i <= 10; i++ {
		fmt.Printf("binary of %d is %[1]b\n", i)
	}

	fmt.Println()
	fmt.Printf("--- \"print octal numbers\" ---\n")
	for i := 0; i <= 10; i++ {
		fmt.Printf("octal of %d is %[1]o\n", i)
	}

	fmt.Println()
	fmt.Printf("--- \"print hexadecimal numbers\" ---\n")
	for i := 0; i <= 20; i++ {
		fmt.Printf("hex of %d is %#[1]X\n", i)
	}
}
