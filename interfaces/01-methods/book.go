package main

import "fmt"

/*
	We cannot have multiple functions or multiple methods with
  same name within the same package
*/

type book struct {
	title string
	price float64
}

// # 1
// func printBook(b book) {
// here b is a copy of the original `book` value
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }

// #2
// func (b book) printBook() {
// 	// here b is still a copy of the original `book` value
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }

// #3
func (b book) print() {
	// b is a copy of the original `book` value here.
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}
