package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("Slice capaity...")
	s.PrintElementAddr = true
	s.PrintBacking = true

	// create a new fruits slice (its a nil slice)
	// the nil slice has no backing array
	var fruits []string
	s.Show("fruits before initialization", fruits)

	fruits = []string{}
	s.Show("fruits overridden with empty slice", fruits)

	// fill values
	fruits = append(fruits, []string{"mangoes", "grapes", "pineapples", "cherries"}...)
	s.Show("fruits after initialization", fruits)
	fmt.Printf("address of fruits: %p\n", &fruits)

	// the size of a slice is always fixed:
	// 24 bytes (on a 64-bit system)
	// slice value = slice header
	fmt.Println(unsafe.Sizeof(""))
	fmt.Printf("size of fruit slice: %d bytes\n", unsafe.Sizeof(fruits))

	part := fruits
	s.Show("part slice", part)
	fmt.Printf("size of part slice: %d bytes\n", unsafe.Sizeof(part))

	// empty part slice
	part = fruits[:0]
	// resize upto 0 element
	s.Show("part[:0]", part)
	// resize upto the capacity
	s.Show("part[:cap]", part[:cap(part)])

	// depleting the capacity of the slice in a loop
	for cap(part) != 0 {
		fmt.Println("performing resize from", cap(part), "to", cap(part)-1)
		part = part[1:cap(part)]
		s.Show("part", part)
	}

	// the games slice should still be ok
	s.Show("fruits slice", fruits)

	// A slice literal creates a new slice value with equal length and capacity.
	s.Show("slice literal", []string{"i", "have", "a", "great", "capacity"})

	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	words = words[2 : cap(words)-2]
	s.Show("www", words)
}
