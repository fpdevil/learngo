package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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

	for i, file := range files {
		if file.Size() != 0 {
			fmt.Printf("%d. %s\n", i, file.Name())
		}
	}
}
