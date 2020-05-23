package problems

import (
	"fmt"
	"time"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

// Exercise8 - Append #2
func Exercise8() {
	fmt.Println("")
	fmt.Println("---- 16-slices/08-append-2 Exercise: Append #2 ----")

	pizza := []string{}
	departures := []time.Time{}
	graduations := []int{}
	lights := []bool{}

	pizza = append(pizza, "jalapeno", "roma tomatoes", "pineapples", "bell pepper")

	// for departure times
	now := time.Now()
	afterOneDay := now.Add(time.Hour * 24)
	afterTwoDays := now.Add(time.Hour * 24 * 2)
	departures = append(departures, now, afterOneDay, afterTwoDays)

	// graduation years
	graduations = append(graduations, []int{1998, 2005, 2018}...)

	// for lights
	lights = append(lights, []bool{true, false, true}...)

	// printing all slices
	fmt.Printf("Pizza toppings: 				%v\n", pizza)
	fmt.Printf("Departure times: 				%v\n", departures)
	fmt.Printf("Student graduation years: 		%v\n", graduations)
	fmt.Printf("On/off states of lights in a room: %v\n", lights)

	fmt.Println()
	fmt.Println("pretty printing the slices...")
	s.Show("Pizza toppings", pizza)
	s.Show("Departure times", departures)
	s.Show("Student graduation years", graduations)
	s.Show("On/off states of lights in a room", lights)
}
