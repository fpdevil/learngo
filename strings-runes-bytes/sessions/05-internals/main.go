package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== string internals ===")

	// string headers of am empty string
	fmt.Println("* String Headers for empty string *")
	empty := ""
	dump(empty)

	fmt.Println("\n\n")
	fmt.Println("* String Headers for hello string *")
	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")

	fmt.Println("\n\n")
	for i := range hello {
		dump(hello[i : i+1])
	}

	fmt.Println("\n\n")
	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello)))
}

// StringHeader is used by a string value
// In practice, you should use: reflect.Header
type StringHeader struct {
	// points to a backing array's item
	pointer uintptr // where it starts
	length  int     // where it ends
}

// dump prints the string header of a string value
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
