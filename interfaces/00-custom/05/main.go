package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("---- Accept Interfaces & Return Structs ----")
	fmt.Println()
}

// loadPersonx accepts a concrete string and returns a struct
func loadPersonx(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)

	if err != nil {
		return p, err
	}
	return p, nil
}

// loadPersony accepts an io.Redaer interface
func loadPersony(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)

	if err != nil {
		return p, err
	}
	return p, nil
}
