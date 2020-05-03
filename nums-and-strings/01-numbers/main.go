// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fpdevil/learngo/nums-and-strings/01-numbers/exercises"
)

//------------------------------------
// Running:
// go run main.go sphere 10
// go run main.go
//------------------------------------

func main() {
	exercises.Exercise1()
	exercises.Exercise2()
	exercises.Exercise3()
	exercises.Exercise4()
	exercises.Exercise5()
	exercises.Exercise6()
	exercises.Exercise7()

	if os.Args != nil && len(os.Args) > 1 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		if arg1 == "sphere" {
			radius, _ := strconv.ParseFloat(arg2, 64)
			exercises.Exercise8(radius)
			exercises.Exercise9(radius)
		} else {
			arg1 = ""
			fmt.Println("unable to decrypt...")
		}
	}
}
