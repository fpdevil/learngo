package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

func (t triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) Name() string {
	return "Triangle"
}

func (t triangle) String() string {
	return fmt.Sprintf("A Triangle of height %f & base %f\n", t.height, t.base)
}
