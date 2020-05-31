package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("03-bytes-runes-strings-examples")

	s.PrintBacking = true
	s.PrintElementAddr = true

	// a string is a read only byte slice
	// once declared, it cannot change
	str := "Yūgen ☯ 💀"
	s.Show("str", str)

	// we can convert string to byte and update the same
	// a byte slice can be changed
	bytes := []byte(str)
	s.Show("bytes", bytes)

	str = string(bytes)
	s.Show("str", str)

	fmt.Println()
	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))

	// len function counts bytes and not runes
	// so use utf8.RuneCountInString
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Println()
	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()
	fmt.Println("--- bytes of the string ---")
	for i := range str {
		fmt.Printf("str[%2d] = %-2x\n", i, str[i])
	}

	fmt.Println()
	fmt.Println("--- same as above using utf8 encoded values ---")
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	// index junps from 1 to 3 as second element occupies 2 bytes
	// 	str[ 0] = 59           = 'Y'
	// 	str[ 1] = c5 ab        = 'ū'
	// 	str[ 3] = 67           = 'g'

	fmt.Println()
	fmt.Printf("1st byte	: %c\n", str[0])
	fmt.Printf("2nd byte	: %c\n", str[1])
	fmt.Printf("2nd rune	: %s\n", str[1:3])
	fmt.Printf("last rune  	: %s\n", str[len(str)-4:])

	// converting to a rune
	runes := []rune(str)
	fmt.Println()
	fmt.Printf("runes from string: %s\n", string(runes))
	fmt.Printf("size of runes: %d bytes\n", int(unsafe.Sizeof(runes[0])) * len(runes))


	fmt.Println()
	fmt.Println("--- printing runes ---")

	word := "öykü"
	fmt.Printf("%q in runes: %c\n", word, []rune(word))
	fmt.Printf("%q in bytes: % [1]x\n", word)

	fmt.Printf("%s %s\n", word[:2], []byte{word[0], word[1]}) // ö
	fmt.Printf("%c\n", word[2])                               // y
	fmt.Printf("%c\n", word[3])                               // k
	fmt.Printf("%s %s\n", word[4:], []byte{word[4], word[5]}) // ü
}
