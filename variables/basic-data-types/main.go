package main

import "fmt"

func main() {
	fmt.Println("Integer Literals...")
	fmt.Println(-200, -100, 0, 50, 100, 200)
	fmt.Println("FloatingPoint Literals...")
	fmt.Println(-90.99, -50.0, 0.0, 50.0, 99.999999)
	fmt.Println("Boolean Constants")
	fmt.Println(true, false)
	fmt.Println("String literals [utf-8]")
	fmt.Println(
		"",
		"masculine",
		"hi there 星!",
	)
}
