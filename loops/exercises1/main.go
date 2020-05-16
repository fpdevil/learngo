package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpdevil/learngo/loops/exercises1/problems"
)

func main() {
	fmt.Println("==== Running Exercises from 13-loops ====")

	problems.Exercise1()
	problems.Exercise2()

	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex3 minval maxval
	if l > 2 && args[1] == "ex3" {
		problems.Exercise3(strings.Join(args[2:], ","))
	} else {
		problems.Exercise3("")
	}

	// go run main.go ex4 minval maxval
	if l > 2 && args[1] == "ex4" {
		problems.Exercise4(strings.Join(args[2:], ","))
	} else {
		problems.Exercise4("")
	}

	// go run main.go ex5 minval maxval
	if l > 2 && args[1] == "ex5" {
		problems.Exercise5(strings.Join(args[2:], ","))
	} else {
		problems.Exercise5("")
	}

	// press Ctrl + C to get out
	// problems.Exercise6()
	if l > 2 && args[1] == "ex6" {
		problems.Exercise6()
	} else {
		fmt.Println("skipping exercise6 for explicit invocation")
	}

	// go run main.go ex7 6
	if l > 2 && args[1] == "ex7" {
		problems.Exercise7(strings.Join(args[2:], ","))
	} else {
		problems.Exercise7("")
	}

	// go run main.go ex8 <operator> <value>
	// use quotes for operator as in
	// go run main.go ex8 "*" 4
	if l > 2 && args[1] == "ex8" {
		problems.Exercise8(strings.Join(args[2:], ","))
	} else {
		problems.Exercise8("")
	}
}
