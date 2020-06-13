package main

import "fmt"

func main() {
	fmt.Println("--- 26-pointers/01-pointers ---")
	fmt.Println()

	var counter byte = 100
	P := &counter
	V := *P

	fmt.Printf("counter	: %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P	: %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V	: %-16d address: %-16p\n", V, &V)

	// assign a value to V
	V = 200
	fmt.Println()
	fmt.Printf("counter	: %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P	: %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V	: %-16d address: %-16p\n", V, &V)

	// reset the value of V to counters initial value
	V = counter
	counter++
	fmt.Println()
	fmt.Printf("counter	: %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P	: %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V	: %-16d address: %-16p\n", V, &V)

	// changing value through pointer
	*P = 25
	fmt.Println()
	fmt.Printf("counter	: %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P	: %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V	: %-16d address: %-16p\n", V, &V)

}
