package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpdevil/learngo/loops/exercises2/problems"
)

func main() {
	fmt.Println("==== Running Exercises from 13-loops ====")

	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex1 guess
	if l > 2 && args[1] == "ex1" {
		problems.Exercise1(strings.Join(args[2:], ","))
	} else {
		problems.Exercise1("")
	}

	// go run main.go ex2 guess
	if l > 2 && args[1] == "ex2" {
		problems.Exercise2(strings.Join(args[2:], ","))
	} else {
		problems.Exercise2("")
	}

	// go run main.go ex3 guess1 guess2
	if l > 2 && args[1] == "ex3" {
		problems.Exercise3(strings.Join(args[2:], ","))
	} else {
		problems.Exercise3("")
	}

	// go run main.go ex4 -v guess
	if l > 2 && args[1] == "ex4" {
		problems.Exercise4(strings.Join(args[2:], ","))
	} else {
		problems.Exercise4("")
	}

	// go run main.go ex5 guess
	if l > 2 && args[1] == "ex5" {
		problems.Exercise5(strings.Join(args[2:], ","))
	} else {
		problems.Exercise5("")
	}

	// go run main.go ex6 guess
	if l > 2 && args[1] == "ex5" {
		problems.Exercise5(strings.Join(args[2:], ","))
	} else {
		problems.Exercise5("")
	}

	// go run main.go ex6 guess
	if l > 2 && args[1] == "ex6" {
		problems.Exercise6(strings.Join(args[2:], ","))
	} else {
		problems.Exercise6("")
	}
}
