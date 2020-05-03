package main

import (
	"fmt"
	"os"

	"github.com/fpdevil/learngo/nums-and-strings/02-strings/exercises"
)

func main() {
	fmt.Println("Running exercises from 08-numbers-and-strings/02-strings")

	exercises.Exercise1()
	exercises.Exercise2()

	if os.Args != nil && len(os.Args) > 1 {
		arg1 := os.Args[1]
		var name string
		if arg1 == "ex3" {
			// go run main.go ex3 mickey
			if len(os.Args[2]) != 0 {
				name = os.Args[2]
			}
			exercises.Exercise3(name)
		} else if arg1 == "ex4" {
			// go run main.go ex4 İNANÇ
			if len(os.Args[2]) != 0 {
				name = os.Args[2]
			}
			exercises.Exercise4(name)
		} else if arg1 == "ex5" {
			// go run main.go ex5 İNANÇ
			if len(os.Args[2]) != 0 {
				name = os.Args[2]
			}
			exercises.Exercise5(name)
		} else if arg1 == "ex6" {
			// go run main.go ex6 MICKEY
			if len(os.Args[2]) != 0 {
				name = os.Args[2]
			}
			exercises.Exercise6(name)
		} else {
			arg1 = ""
			fmt.Println("unable to decrypt...")
		}
	}

	exercises.Exercise7()
	exercises.Exercise8()
}
