package main

import "fmt"

// Shape interface
type Shape interface {
	Area() float64
	Name() string
}

func displayShapes(shapes ...Shape) {
	for _, v := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", v.Name(), v.Area())
	}
}
