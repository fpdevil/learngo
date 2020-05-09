package problems

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "inanc"
	pass, pass2 = "1888", "1879"
)

// Exercise3 - Convert
func Exercise3(xargs string) {
	fmt.Println("")
	fmt.Println("---- 12-switch/03-convert Exercise: Convert ----")

	args := strings.Split(xargs, ",")
	// fmt.Println("argument string: ", args, "length: ", len(args))

	// args := os.Args

	// if len(args) != 3 {
	// 	fmt.Println(usage)
	// 	return
	// }

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	// u, p := args[1], args[2]
	u, p := args[0], args[1]

	//
	// REFACTOR THIS TO A SWITCH
	//
	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }

	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		fmt.Printf(accessOK, u)
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
