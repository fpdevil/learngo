package main

import "fmt"

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.base * r.height
}

func (r rectangle) Name() string {
	return "Rectangle"
}

func (r rectangle) String() string {
	return fmt.Sprintf("A Rectangle of height %f & base %f\n", r.height, r.base)
}
