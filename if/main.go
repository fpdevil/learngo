package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpdevil/learngo/if/exercises"
)

func main() {
	fmt.Println("==== Running Exercises from 11-if ====")

	exercises.Exercise1()
	exercises.Exercise2()

	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex3 hh jj kk
	if l > 2 && args[1] == "ex3" {
		exercises.Exercise3(strings.Join(args[2:], ","))
	} else {
		exercises.Exercise3("")
	}

	// go run main.go ex4 a
	if l > 2 && args[1] == "ex4" {
		exercises.Exercise4(strings.Join(args[2:], ","))
	} else {
		exercises.Exercise4("")
	}

	// go run main.go ex5 16
	if l > 2 && args[1] == "ex5" {
		exercises.Exercise5(strings.Join(args[2:], ","))
	} else {
		exercises.Exercise5("")
	}

	// go run main.go ex6 24
	if l > 2 && args[1] == "ex6" {
		exercises.Exercise6(strings.Join(args[2:], ","))
	} else {
		exercises.Exercise6("")
	}

	// go run main.go ex7 2019
	if l > 2 && args[1] == "ex7" {
		exercises.Exercise7(strings.Join(args[2:], ","))
	} else {
		exercises.Exercise7("")
	}

	// go run main.go ex9 "MaRcH"
	if l > 2 && args[1] == "ex9" {
		exercises.Exercise9(strings.Join(args[2:], ","))
	} else {
		exercises.Exercise9("")
	}
}
