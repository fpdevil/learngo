package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		n, s := fizzBuzz(i)
		fmt.Printf("results: %d %s\n", n, s)
	}
}

func fizzBuzz(num int) (int, string) {
	switch  {
	case num % 15 == 0:
		return num, "FizzBuzz"
	case num % 3 == 0:
		return num, "Fizz"
	case num % 5 == 0:
		return num, "Buzz"
	}
	return num, ""
}
