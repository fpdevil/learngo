package main

import "fmt"

type cat struct {
	name string
	age  int
}

func (c cat) Speak() {
	fmt.Printf("%-15s %-2d\n", c.name, c.age)
}

func (c cat) String() string {
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}
