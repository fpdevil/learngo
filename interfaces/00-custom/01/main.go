package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing interfaces...")

	var (
		c = cat{name: "Thomas the Cat", age: 5}
		d = dog{name: "Spike", age: 6}
		r = rocket{model: "saturn V5"}
	)

	var z list
	z = append(z, &c, &d, &r)

	z.speaks()

	// c.Speak()
	items := []Speaker{cat{"meow", 2}, dog{"woof woof", 3}, &rocket{"ariane"}}
	for _, i := range items {
		i.Speak()
	}

	fmt.Println(r)
}
