package main

import "fmt"

func main() {
	fmt.Println("shapes and details...")
	fmt.Println()

	var (
		t = triangle{base: 21.2, height: 36.3}
		r = rectangle{base: 11.0, height: 6.0}
		s = square{side: 12.12}
	)

	displayShapes(t, r, s)

	fmt.Println()
	fmt.Println(t)
	fmt.Println(r)
	fmt.Println(s)
}
