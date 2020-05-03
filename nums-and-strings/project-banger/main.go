package main

import (
	"fmt"
	"os"
	"strings"
)

//------------------------
// get input + bang it
// NOTE: You should always pass it at least one argument
// eg: go run main.go flask
// returns flask!!!!!
//
// EXERCISE: prepend with the bang too
//------------------------

func main() {
	msg := os.Args[1]

	// The len built-in function returns the length of v, according to its type:
	//	for String: the number of bytes in v.
	l := len(msg)

	// for string length, we have to use utf8.RuneCountInString
	bangs := strings.Repeat("!", l)
	s := bangs + msg + bangs
	s = strings.ToUpper(s)

	fmt.Println(s)
}
