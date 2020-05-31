package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("02-english-dict-map-populate")
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	fmt.Println()
	query := args[0]

	// an empty map literal
	// dict := map[string][string]
	// fmt.Printf("Zero valueL %#v\n", dict)

	// defining a map literal
	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	dict["up"] = "yukarı"  // adds new kv
	dict["down"] = "aşağı" // adds new kv
	dict["good"] = "iyi"   // this updates or overwrites kv
	dict["mistake"] = ""   // key with a zero value

	// if the value if found for a key return
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	// else
	fmt.Printf("%q not found\n", query)

	fmt.Printf("no. of keys: %d\n", len(dict))

	// map maually copied to demonstrate comparision
	copied := map[string]string{"up": "yukarı", "down": "aşağı",
		"mistake": "", "good": "iyi", "great": "harika",
		"perfect": "mükemmel"}
	first := fmt.Sprintf("%s", dict)
	second := fmt.Sprintf("%s", copied)

	fmt.Println("comparing first and second maps...")
	if first == second {
		fmt.Println("maps are equal")
	}

	// print the map indexed output
	// ordered output is available only from Go 1.12
	fmt.Printf("%#v\n", dict)

	// loop over the key value pairs of dict
	// this is not efficient
	fmt.Println("loop over the key value pairs of dict...")
	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v)
	}
}
