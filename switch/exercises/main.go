package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpdevil/learngo/switch/exercises/problems"
)

func main() {
	fmt.Println("==== Running Exercises from 12-switch ====")

	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex1 2.5
	if l > 2 && args[1] == "ex1" {

		problems.Exercise1(strings.Join(args[2:], ","))
	} else {
		problems.Exercise1("")
	}

	// go run main.go ex2 strong
	if l > 2 && args[1] == "ex2" {

		problems.Exercise2(strings.Join(args[2:], ","))
	} else {
		problems.Exercise2("")
	}

	// go run main.go ex3 user pass
	if l > 2 && args[1] == "ex3" {

		problems.Exercise3(strings.Join(args[2:], ","))
	} else {
		problems.Exercise3("")
	}

	// go run main.go ex4 lower 'OMG!'
	if l > 2 && args[1] == "ex4" {
		problems.Exercise4(strings.Join(args[2:], ","))
	} else {
		problems.Exercise4("")
	}

	// go run main.go ex5 lower 'SePTembeR'
	if l > 2 && args[1] == "ex5" {
		problems.Exercise5(strings.Join(args[2:], ","))
	} else {
		problems.Exercise5("")
	}

	fmt.Println("==== completed 12-switch ====")
}
