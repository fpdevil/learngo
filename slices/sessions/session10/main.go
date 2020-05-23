package main

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	fmt.Println("Mechanics of append...")
	fmt.Println("Growth of Backing Array in Animation")

	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int
	screen.Clear()

	// loop until the capacity of slice is greater than 123
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
		// append a random number between 1 and 9
		nums = append(nums, rand.Intn(9)+1)

		// sleep for a period until the next loop step
		time.Sleep(time.Second / 4)
	}
}
