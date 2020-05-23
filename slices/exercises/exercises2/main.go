package main

import (
	"fmt"

	// import the problems package containing exercises
	"github.com/fpdevil/learngo/slices/exercises/exercises2/problems"
)

// this is the main function for running the exercises under 16-slices
func main() {
	fmt.Println("==== Running Exercises from 16-slices ====")

	// var (
	// 	args = os.Args
	// 	l    = len(args)
	// )

	// // go run main.go ex9 4 6 1 3 0 9 2
	// if l > 2 && args[1] == "ex10" {
	// 	problems.Exercise10(strings.Join(args[2:], " "))
	// } else {
	// 	problems.Exercise10("")
	// }

	problems.Exercise16()
	problems.Exercise17()
	problems.Exercise18()
	problems.Exercise19()
	problems.Exercise20()
	problems.Exercise21()
}
