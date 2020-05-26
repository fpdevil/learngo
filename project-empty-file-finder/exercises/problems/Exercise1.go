package problems

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

// Exercise1 - Sort and write items to a file
func Exercise1(vals string) {
	fmt.Println("")
	fmt.Println("---- 1-sort-to-a-file EXERCISE: Sort and write items to a file ----")

	args := strings.Fields(vals)
	// fmt.Println("argument string: ", args, "length: ", len(args))

	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)
	var slice []byte

	for _, v := range args {
		slice = append(slice, v...)
		slice = append(slice, '\n')
	}

	err := ioutil.WriteFile("out.txt", slice, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
