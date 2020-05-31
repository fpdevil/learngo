package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	// #1 Nil map (read only)
	var dict map[string]string

	// values cannot be assiged to a nil map
	// dict["left"] = "right" --- will not work

	// map average retrieval rate is O(1)
	key := "good"

	value := dict[key]
	fmt.Printf("%q means %#v\n", key, value)

	fmt.Printf("# of Keys: %d\n", len(dict))
	fmt.Printf("Zero Value: %#v\n", dict)

	// #4: Nil map ready to use
	if dict != nil {
		value := dict[key]
		fmt.Printf("%q means %#v\n", key, value)
	}

	// #3: Cannot use non-comparable types as map key types
	// var broken map[[]int]int
	// var broken map[map[int]string]bool
	//
	// A map can only be compared to nil value
	// _ = dict == nil
}
