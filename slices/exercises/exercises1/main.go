package main

import (
	"fmt"

	// import the problems package containing exercises
	"github.com/fpdevil/learngo/slices/exercises/exercises1/problems"
)

// this is the main function for running the exercises under 16-slices
func main() {
	fmt.Println("==== Running Exercises from 16-slices ====")

	problems.Exercise1()
	problems.Exercise2()
	problems.Exercise3()
	problems.Exercise4()
	problems.Exercise5()
	problems.Exercise6()
	problems.Exercise7()

	// var (
	// 	args = os.Args
	// 	l    = len(args)
	// )

	// // go run main.go ex9 10.5
	// if l > 2 && args[1] == "ex9" {
	// 	problems.Exercise9(strings.Join(args[2:], ","))
	// } else {
	// 	problems.Exercise9("")
	// }
}
