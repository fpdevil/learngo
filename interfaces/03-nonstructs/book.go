package main

import "fmt"

type book struct {
	title string
	price money
	// price float64
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}
