package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// find empty files and write them to a file

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Give directory as argument")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// names variable will hold the list of empty files
	var names []byte

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()

			fmt.Println(cap(names))
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	// for file writing
	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", names)
}
