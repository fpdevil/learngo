package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("session3 slices continued...")

	// TODO Tasks
	var todo []string
	todo = append(todo, "sing")
	todo = append(todo, "fight", "dance", "code")

	// appending two slices using ellipsis (...)
	tomorrow := []string{"morning walk", "eat breakfast"}
	todo = append(todo, tomorrow...)

	s.Show("todo", todo)

	x := []int{6, 7, 8}
	y := append(x, []int{11, 12, 13}...)

	fmt.Println(x)
	fmt.Println(y)
}
