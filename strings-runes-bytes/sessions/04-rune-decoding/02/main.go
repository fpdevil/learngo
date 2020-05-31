package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("--- [02] 04-rune-decoding ---")

	// a turkish word in byte slice
	word := []byte("öykü")
	fmt.Printf("%s = % [1]x\n", word)

	// how to make the first rune uppercase?
	// you need to find the starting and ending position of the first rune

	// 1st: Using a `for range`
	// we cannot get the runes by range overing a byte slice
	// we first have to convert to a string
	var size int
	for i := range string(word) {
		if i > 0 {
			size = i
			break
		}
	}

	// 2nd: By using the utf8 package's DecodeRune function
	_, size = utf8.DecodeRune(word)

	// overwrite the current bytes with the new uppercased bytes
	copy(word[:size], bytes.ToUpper(word[:size]))

	// to get printed bytes/runes need to be encoded in a string
	fmt.Printf("%s = % [1]x\n", word)

}
