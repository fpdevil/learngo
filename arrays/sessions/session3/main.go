package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly #2
//
//   This challenge is based on the previous Moodly challenge.
//
//   1. Display the usage if the username or the mood is missing
//
//   2. Change the moods array to a multi-dimensional array.
//
//      So, create two inner arrays:
//        1. One for positive moods
//        2. Another one for negative moods
//
//   4. Randomly select and print one of the mood messages depending
//      on the given mood command-line argument.
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name] [positive|negative]
//
//   go run main.go Socrates
//     [your name] [positive|negative]
//
//   go run main.go Socrates positive
//     Socrates feels good 👍
//
//   go run main.go Socrates positive
//     Socrates feels happy 😀
//
//   go run main.go Socrates positive
//     Socrates feels awesome 😎
//
//   go run main.go Socrates negative
//     Socrates feels bad 👎
//
//   go run main.go Socrates negative
//     Socrates feels sad 😞
//
//   go run main.go Socrates negative
//     Socrates feels terrible 😩
// ---------------------------------------------------------

const usageMsg = `
	[your name] [positive|negative]
`

func main() {
	// multi-dimensional arrays
	// student1 := [3]int{5, 6, 1}
	// student2 := [3]int{9, 8, 4}
	// students := [2][3]float64{
	// 	[3]float64{5, 6, 1},
	// 	[3]float64{9, 8, 4},
	// }
	students := [...][3]float64{
		{5, 6, 1},
		{9, 8, 4},
	}

	var sum float64
	// sum += students[0][0] + students[0][1] + students[0][2]
	// sum += students[1][0] + students[1][1] + students[1][2]
	for _, grades := range students {
		for _, grade := range grades {
			sum += grade
		}
	}

	const N = float64(len(students) * len(students[0]))
	fmt.Printf("Average Grade: %g\n", sum/N)

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(usageMsg)
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	t := time.Now()
	rand.Seed(t.UnixNano())
	limit := len(moods[0])
	index := rand.Intn(limit)

	// index for positive or negative
	var pni int

	var usermood string
	// fmt.Println(index)
	if mood != "positive" {
		pni = 1
	}
	usermood = moods[pni][index]

	fmt.Printf("%s feels %s\n", name, usermood)
}
