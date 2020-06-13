package main

import (
	"fmt"
)

type cat struct {
	name string
}

func main() {
	fmt.Println("--- interface types ---")
	fmt.Println()

	c := cat{name: "figaro"}
	i := []interface{}{121, "The Book of Enoch", true, c, `a`}
	getTypes(i)
}

func getTypes(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is INT\n", v)
		case string:
			fmt.Printf("%v is STRING\n", v)
		case bool:
			fmt.Printf("%v is BOOL\n", v)
		default:
			fmt.Printf("%v is unknown of type %[1]T\n", v)
		}
	}
}
