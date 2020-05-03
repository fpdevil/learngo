package main

import "fmt"

// running => go run *.go
func main() {
	fmt.Println("Hello Gophers!")
	// You can access functions from other files
	// which are in the same package
	bye()
	hey()
}
