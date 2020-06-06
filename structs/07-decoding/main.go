package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// json decoding
// ----------------------------

// user struct for decoding and holding the user data
// using field tags
type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

// this gives response [{ map[]} { map[]} { map[]}] as the
// user variables from json is different than what we have
// defined here.
// type user struct {
// 	Name        string
// 	Permissions map[string]bool
// }

func main() {
	fmt.Println("--- 24-structs/07-decoding ---")

	// declare a new byte slice for the input
	var input []byte

	// define a scanner for reading the input from stdin
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Bytes()...)

		// appending a newline for output rendering/pretty printing
		// input = append(input, '\n')
	}

	// fmt.Printf("%s\n", string(input))

	// decoding the json
	var users []user

	// this fails with json: Unmarshal(non-pointer []main.user)
	// reason: the function cannot change the passed in parameter
	// users and write the output to it.
	// for that we have to pass a pointer to users
	// err := json.Unmarshal(input, users)

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// fmt.Printf("%v\n", users)
	// this prints the below if proper field tags are not used
	// [{ map[]} { map[]} { map[]}]

	fmt.Printf("%#v\n", users)

	// pretty printing the output struct
	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permissions; {
		case p == nil:
			fmt.Println(" user has no power.")
		case p["admin"]:
			fmt.Println(" user is an admin.")
		case p["write"]:
			fmt.Println(" user can write.")
		}
		fmt.Println()
	}
}
