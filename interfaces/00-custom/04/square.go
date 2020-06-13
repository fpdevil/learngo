package main

import "fmt"

type square struct {
	side float64
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "Square"
}

func (s square) String() string {
	return fmt.Sprintf("A Squareof sides %f X %[1]f\n", s.side)
}
