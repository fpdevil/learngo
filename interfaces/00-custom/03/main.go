package main

import "fmt"

type speaker interface {
	speak() string
}

type cat struct{}

type dog struct{}

type person struct {
	name string
}

func main() {
	fmt.Println("interfaces and concrete types...")
	fmt.Println()

	var (
		c = cat{}
		d = dog{}
		p = person{name: "Hamilton"}
	)

	catSpeaks(c)
	dogSpeaks(d)
	personSpeaks(p)

	speakSomething(c, d, p)
}

func (c cat) speak() string {
	return "Meow Meow Purrr..."
}

func (d dog) speak() string {
	return "Woof woof Grrr..."
}

func (p person) speak() string {
	return "Hello, from " + p.name + "!"
}

func catSpeaks(c cat) {
	fmt.Println(c.speak())
}

func dogSpeaks(d dog) {
	fmt.Println(d.speak())
}

func personSpeaks(p person) {
	fmt.Println(p.speak())
}

func speakSomething(say ...speaker) {
	for _, s := range say {
		fmt.Println(s.speak())
	}
}
