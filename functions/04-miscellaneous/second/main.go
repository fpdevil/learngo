package main

import "fmt"

func main() {
	fmt.Println("--- closures ---")

	calculator(add, 5, 6)
	calculator(subtract, 10, 6)

	v := square(9)
	fmt.Printf("square of 9: %d and type : %T\n", v(), v)
}

func calculator(f func(int, int) int, i, j int) {
	fmt.Println(f(i, j))
}

func add(i, j int) int {
	return i + j
}

func subtract(i, j int) int {
	return i - j
}

func square(x int) func() int {
	f := func() int {
		return x * x
	}
	return f
}
