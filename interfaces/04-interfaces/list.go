package main

import "fmt"

// create an interface for common printing of both the
// books and games
// an abstract type, a protocol or a contract with out
// a concrete implementation
type printer interface {
	print() // only describes print behaviour
}

// type list []*game

// func (l list) print() {
// 	if len(l) == 0 {
// 		fmt.Println("Sorry. We're waiting for delivery 🚚.")
// 		return
// 	}

// 	for _, it := range l {
// 		it.print()
// 	}
// }

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		fmt.Printf("{ %-12T }  ===> ", it)
		it.print()
	}
}
