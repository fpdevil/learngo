package main

import (
	"fmt"
	"strconv"
)

// N is a package level variable
// it belongs to the main package
// var N int

func main() {
	local := 10
	show(local)

	incrWrong(local)
	fmt.Printf("local + %d\n", local)

	// incr(local)
	// fmt.Printf("now local + %d\n", local)
	local = incr(local)
	show(local)

	local = incrBy(local, 5)
	show(local)

	local, _ = incrByStr(local, "2")
	show(local)

	// function chaining
	// we cannot save the number back into mains local
	show(incrBy(local, 2))
	show(local)

	local = sanitize(incrByStr(local, "2"))
	show(local)

	local = limit(incrBy(local, 5), 2000)
	show(local)
}

func show(n int) {
	// local defined in main cannot be accessed
	// fmt.Printf("show n = %d\n", local)
	fmt.Printf("show n = %d\n", n)
}

// this incrWrong can't access to main's `local`
func incrWrong(n int) {
	// n := local
	n++
}

func incr(n int) int {
	n++
	return n
}

func incrBy(n, factor int) int {
	return n + factor
}

func incrByStr(n int, factor string) (int, error) {
	m, err := strconv.Atoi(factor)
	n = incrBy(n, m)
	return n, err
}

func sanitize(n int, err error) int {
	if err != nil {
		return 0
	}
	return n
}

func limit(n, lim int) (m int) {
	// var m int
	m = n
	if m >= lim {
		m = lim
	}
	// return m
	return
}
