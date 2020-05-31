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

	query := args[0]

	english := []string{"good", "great", "perfect"}
	turkish := []string{"iyi", "harika", "mükemmel"}

	for i, w := range english {
		if query == w {
			fmt.Printf("%q means %q in turkish\n", w, turkish[i])
			return
		}
	}
	fmt.Println("Rate is O(n)")
	fmt.Printf("%q in english not found\n", query)
}
