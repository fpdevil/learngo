package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fpdevil/learngo/arrays/exercises/exercises1/problems"
)

func main() {
	fmt.Println("==== Running Exercises from 14-arrays ====")

	problems.Exercise1()
	problems.Exercise3()
	problems.Exercise5()
	problems.Exercise7()
	problems.Exercise8()

	var (
		args = os.Args
		l    = len(args)
	)

	// go run main.go ex9 10.5
	if l > 2 && args[1] == "ex9" {
		problems.Exercise9(strings.Join(args[2:], ","))
	} else {
		problems.Exercise9("")
	}

	// go run main.go ex10 sTaY
	if l > 2 && args[1] == "ex10" {
		problems.Exercise10(strings.Join(args[2:], ","))
	} else {
		problems.Exercise10("")
	}

	// go run main.go ex11 1 a 2 b 3
	if l > 2 && args[1] == "ex11" {
		problems.Exercise11(strings.Join(args[2:], ","))
	} else {
		problems.Exercise11("")
	}

	// go run main.go ex12 5 4 a c 1
	if l > 2 && args[1] == "ex12" {
		problems.Exercise12(strings.Join(args[2:], ","))
	} else {
		problems.Exercise12("")
	}

	// go run main.go ex13 Cat Beginning
	if l > 2 && args[1] == "ex13" {
		problems.Exercise13(strings.Join(args[2:], ","))
	} else {
		problems.Exercise13("")
	}
}
