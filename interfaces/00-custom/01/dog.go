package main

import "fmt"

type dog struct{
	name string
	age int
}

func (d dog) Speak() {
	fmt.Printf("%-15s %-2d\n", d.name, d.age)
}

func (d dog) String() string {
	return fmt.Sprintf("%v (%v years old)", d.name, d.age)
}
