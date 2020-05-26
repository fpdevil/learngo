package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpdevil/learngo/project-empty-file-finder/exercises/problems"
	// import the problems package containing exercises
)

// this is the main function for running the exercises under 16-slices
func main() {
	fmt.Println("==== Running Exercises from 17-project-empty-file-finder ====")

	// handle the command line arguments
	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex1 orange banana apple
	if l > 2 && args[1] == "ex1" {
		problems.Exercise1(strings.Join(args[2:], " "))
	} else {
		problems.Exercise1("")
	}

	// go run main.go ex2 orange banana apple
	if l > 2 && args[1] == "ex2" {
		problems.Exercise2(strings.Join(args[2:], " "))
	} else {
		problems.Exercise2("")
	}

	// go run main.go ex3 dir/ dir2/
	if l > 2 && args[1] == "ex3" {
		problems.Exercise3(strings.Join(args[2:], " "))
	} else {
		problems.Exercise3("")
	}
}
