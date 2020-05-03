package main

import (
	"fmt"

	"github.com/fpdevil/learngo/first/printer"
)

func main() {
	printer.Hello()
	// for printer exercise
	fmt.Println("Go Version is: " + printer.GoVersion())
}
