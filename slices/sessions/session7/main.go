package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

// type collection [4]string
type collection []string

func main() {
	fmt.Println("Slice Headers...")
	s.PrintElementAddr = true
	data := collection{"mangoes", "grapes", "pineapples", "cherries"}

	change(data)
	s.Show("Mains data", data)
	fmt.Printf("address of mains data: %p\n", &data)

	array := [...]string{"mangoes", "grapes", "pineapples", "cherries"}

	// the size of an array depends on its elements
	fmt.Printf("size of array: %d bytes\n", unsafe.Sizeof(array))

	// the size of a slice is always fixed:
	// 24 bytes (on a 64-bit system)
	// slice value = slice header
	fmt.Printf("size of slice: %d bytes\n", unsafe.Sizeof(data))
}

func change(data collection) {
	data[2] = "blue berries"
	s.Show("changed data", data)
	fmt.Printf("address of changes data: %p\n", &data)
}
