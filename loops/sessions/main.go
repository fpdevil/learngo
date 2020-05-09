package main

import (
	"fmt"
	"strings"
)

func main() {

	const max = 5

	var (
		sum int
		i   = 1
	)

	// insert a debugger here to test
	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			// will fail without i++
			// this continues foever as the number i is never
			// incremented and always starts again from base value
			// to fix increment i as below
			i++
			continue
		}

		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)

	fmt.Println("*** 	Multiplication Table 	***")
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}

	fmt.Println("/ 	Looping over Slices/Words 	/")
	// this will split the following string by spaces
	// and then it will put it inside words variable
	// as a slice of strings
	words := strings.Fields("lazy cat jumps again and again and again")

	// the above will give
	// [lazy cat jumps again and again and again]
	fmt.Println("/ 	using for over words 	/")
	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}

	// ranges
	// range returns next index and value
	fmt.Println("/ 	using range over words 	/")
	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
}
