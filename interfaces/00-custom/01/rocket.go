package main

import "fmt"

type rocket struct {
	model string
}

func (r *rocket) Speak() {
	fmt.Printf("%-15s\n", r.model)
}

func (r rocket) String() string {
	return fmt.Sprintf("The rocket model is %v", r.model)
}

