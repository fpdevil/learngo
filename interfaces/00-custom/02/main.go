package main

import "fmt"


func main() {
	fmt.Println("custom interfaces....")
	fmt.Println()

	fruits := []string{"apples", "grapes", "mangoes", "oranges", "cherries"}

	// this gives the below error
	// cannot use fruits (type []string) as type []interface {} in argument to showAll
	// showAll(fruits)

	vals := make([]interface{}, len(fruits))
	for i, v := range fruits {
		vals[i] = v
	}
	showAll(vals)
}

func showAll(vals []interface{}) {
	for _, v := range vals {
		fmt.Println(v)
	}
}
