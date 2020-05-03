package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("greeter")

	// var name string

	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Number of people: ", len(os.Args)-1)
	name1 := os.Args[1]
	name2 := os.Args[2]
	name3 := os.Args[3]

	fmt.Println("Hello ", name1, "!")
	fmt.Println("Hello ", name2, "!!")
	fmt.Println("Hey ", name3, "!!")

	// fmt.Println("Path: ", os.Args[0])
	// fmt.Println("1st Argument:", os.Args[1])
	// fmt.Println("2nd Argument:", os.Args[2])
	// fmt.Println("3rd Argument:", os.Args[3])

}
