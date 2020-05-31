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
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("---- 22-maps/03-students ----")
	fmt.Println()

	// House        Student Name
	// ---------------------------
	// gryffindor   weasley
	// gryffindor   hagrid
	// gryffindor   dumbledore
	// gryffindor   lupin
	// hufflepuf    wenlock
	// hufflepuf    scamander
	// hufflepuf    helga
	// hufflepuf    diggory
	// ravenclaw    flitwick
	// ravenclaw    bagnold
	// ravenclaw    wildsmith
	// ravenclaw    montmorency
	// slytherin    horace
	// slytherin    nigellus
	// slytherin    higgs
	// slytherin    scorpius
	// bobo         wizardry
	// bobo         unwanted

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	// define the hogwarts map and
	// populate data into hogwarts map
	hogwarts := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// keys := reflect.ValueOf(hogwarts).MapKeys()
	delete(hogwarts, "bobo")

	house := args[0]
	students := hogwarts[house]
	if len(students) == 0 {
		fmt.Printf("Sorry. I don't know anything about %q.\n", house)
		return
	}

	//  "bobo" doesn't belong to Hogwarts, remove it by using
	//  the delete function.

	list := make([]string, 0)
	list = append(list, students...)
	sort.Strings(list)

	fmt.Printf("~~~ %s students ~~~\n\n", house)
	for _, v := range list {
		fmt.Printf("\t+ %s\n", v)
	}
}
