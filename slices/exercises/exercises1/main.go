package main

import (
	"fmt"
	"os"
	"strings"

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
	problems.Exercise8()
	problems.Exercise9()

	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex9 4 6 1 3 0 9 2
	if l > 2 && args[1] == "ex10" {
		problems.Exercise10(strings.Join(args[2:], " "))
	} else {
		problems.Exercise10("")
	}

	problems.Exercise11()
	problems.Exercise12()
	problems.Exercise13()

	// go run main.go ex14 1 2 4
	if l > 2 && args[1] == "ex14" {
		problems.Exercise14(strings.Join(args[2:], " "))
	} else {
		problems.Exercise14("")
	}

	// go run main.go ex15 Location
	if l > 2 && args[1] == "ex15" {
		problems.Exercise15(strings.Join(args[2:], " "))
	} else {
		problems.Exercise15("")
	}
}
