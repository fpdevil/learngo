package main

import (
	"fmt"
	"os"

	"github.com/fpdevil/learngo/go-type-system/exercises/lessons"
)

// EXERCISES FROM 09-go-type-system
func main() {
	fmt.Println("Running exercises from 09-go-type-system")

	lessons.Exercise1()
	lessons.Exercise2()

	if os.Args != nil && len(os.Args) > 1 {
		arg1 := os.Args[1]

		if arg1 == "ex3" {
			// go run main.go ex3 50 25000 2000000 50000000000000000 00000100
			if len(os.Args) == 7 {
				lessons.Exercise3(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
			} else {
				lessons.Exercise3("50", "25000", "2000000", "50000000000000000", "00000100")
			}
		} else if arg1 == "ex4" {
			// go run main.go ex4 2
			if len(os.Args) == 3 {
				lessons.Exercise4(os.Args[2])
			} else {
				lessons.Exercise4("2")
			}
		} else if arg1 == "ex5" {
			// go run main.go ex5 3
			if len(os.Args) == 3 {
				lessons.Exercise5(os.Args[2])
			} else {
				lessons.Exercise5("4")
			}
		} else {
			arg1 = ""
			fmt.Println("unable to decrypt...")
		}
	} else {
		lessons.Exercise3("50", "25000", "2000000", "50000000000000000", "00000100")
		lessons.Exercise4("2")
		lessons.Exercise5("4")
	}

	lessons.Exercise6()
}
