package exercises

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

// Exercise4 - Vowel or Consonant
func Exercise4(vow string) {
	fmt.Println("")
	fmt.Println("---- 11-if/04-vowel-or-cons Exercise: Vowel or Consonant ----")

	vows := strings.Split(vow, ",")
	// fmt.Printf("%#v, %T, %d\n", vows, vows, len(vows))
	if vows[0] == "" {
		fmt.Println("Give me a letter")
		return
	}

	s := vows[0]
	if strings.IndexAny(s, "aeiou") != -1 {
		fmt.Printf("%q is vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}
