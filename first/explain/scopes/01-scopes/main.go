package main

// at file scope
import "fmt"

// at package scope
const ok = true

// at package scope
func main() { //  block scope
	var hello = "Hello!"
	// hello and ok are visible here
	fmt.Println(hello, ok)
}
