package main

import "fmt"

type huge struct {
	// a huge value about ~200mb
	games [10000000]game
}

// get address if the receiver
// calling addr() x 10 times = ~2 GB.
// each time the method is called,
// the original value will be copied.
// func (h huge) addr() {
// 	fmt.Printf("%p\n", &h)
// }

// changing to pointer
func (h *huge) addr() {
	fmt.Printf("%p\n", h)
}
