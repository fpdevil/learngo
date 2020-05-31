package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("bytes-runes-strings-basics")
	str := "hey"
	bytes := []byte{104, 101, 121}

	fmt.Printf(`"hey" in bytes: %d`+"\n", []byte(str))
	fmt.Printf("bytes in string: %q\n", string(bytes))

	// runes are unicode codepoints (numbers)
	fmt.Println()
	fmt.Printf("%c			: %[1]d\n", 'h')
	fmt.Printf("%c			: %[1]d\n", 'e')
	fmt.Printf("%c			: %[1]d\n", 'y')

	// a rune literal is typeless
	// you can put it in any numeric type
	var (
		anInt   int   = 'h'
		anInt8  int8  = 'h'
		anInt16 int16 = 'h'
		anInt32 int32 = 'h'

		// rune literal's default type is: rune, so we don't have to speciify it
		aRune = 'h'
	)

	fmt.Println()
	fmt.Printf("Rune literals are typeless: { %T, %T, %T, %T, %T }\n", anInt, anInt8, anInt16, anInt32, aRune)
	fmt.Println()

	fmt.Printf("%q in decimal: %[1]d\n", 104)
	fmt.Printf("%q in binary: %08[1]b\n", 'h')
	fmt.Printf("%q in binary: 0x%[1]x\n", 0x68)

	fmt.Println()
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	s := "éक्षिaπ囧"
	for i, r := range s {
		fmt.Printf("%2d: %-12v 0x%-12[2]x % -12x %-12[3]q\n", i, r, string(r))
	}
}
