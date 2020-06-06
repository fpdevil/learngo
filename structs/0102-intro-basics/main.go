package main

import "fmt"

func main() {
	fmt.Println("---- 24-structs/0102-intro-basics ----")
	fmt.Println("\n")

	/*
		USING VARIABLES
	*/
	var (
		name, lastname string
		age            int

		name2, lastname2 string
		age2             int
	)

	name, lastname, age = "Pablo", "Picasso", 95
	name2, lastname2, age = "Sigmund", "Freud", 83

	fmt.Println("Picasso:", name, lastname, age)
	fmt.Println("Freud  :", name2, lastname2, age2)

	// This is an anonymous struct as it does not have a name yet
	// we have to duplicate the same struct type for declaring again
	var picasso struct {
		name, lastname string
		age            int
	}

	var freud struct {
		name, lastname string
		age            int
	}

	fmt.Printf("type => %T\n", picasso)
	fmt.Printf("Picasso struct %+v\n", picasso)
	fmt.Printf("Sigmun Freud struct %+v\n", freud)

	picasso.name = "Pablo"
	picasso.lastname = "Picasso"
	picasso.age = 36

	freud.name = "Freud"
	freud.lastname = "Sigmund"
	freud.age = 41

	fmt.Printf("Picasso  %+v\n", picasso)
	fmt.Printf("Sigmun Freud  %+v\n", freud)

	// New struct type declarations
	type Person struct {
		Name     string
		Lastname string
		Age      int
	}

	var donald Person
	donald.Name = "Donald"
	donald.Lastname = "Duck"
	donald.Age = 121
	fmt.Printf("Donald Duck: %#v\n", donald)

	person1 := Person{Name: "Mickey", Lastname: "Mouse", Age: 99}
	person2 := Person{Name: "Gyro", Lastname: "Gearloose", Age: 100}

	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("person2: %+v\n", person2)

	type VideoGame struct {
		Title, Genre string
		Published    bool
	}

	pokemon := VideoGame{
		Title:     "Pokemon",
		Genre:     "Game and Cartoon",
		Published: true,
	}

	fmt.Printf("pokemon: %#v\n", pokemon)

	type movie struct {
		title, genre string
		rating       float64
		released     bool
	}

	avengers := movie{"avengers: end game", "sci-fi", 8.9, true}
	clone := movie{
		title: "avengers: end game", genre: "sci-fi",
		rating: 8.9, released: true,
	}

	fmt.Println(avengers)
	fmt.Println(clone)
	fmt.Print("avengers == clone?\n", avengers == clone)
}
