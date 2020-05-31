// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 0 from Product ID #576872813.
//
// ---------------------------------------------------------

func main() {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("---- 22-maps/02-populate ----")

	// #1
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}

	// #2
	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}

	// #3
	mulphones := map[string][]string{
		"bowen": []string{"202-555-0179"},
		"dulin": []string{"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": []string{"03489940240", "03489900120"},
	}

	// #3
	basket := map[int]map[int]int{}
	basket[100] = map[int]int{617841573: 4, 576872813: 2}
	basket[101] = map[int]int{576872813: 5, 657473833: 20}
	basket[102] = map[int]int{}

	// for case #1
	//     Print the dulin's phone number.
	lname, phnum := "dulin", "N/A"

	if v, ok := phones[lname]; ok {
		phnum = v
	}
	fmt.Printf("%s's phone number: %s\n", lname, phnum)

	// for case #2
	//     Is Product ID 879401371 available?
	pid, status := 879401371, "available"
	if !products[pid] {
		status = "not " + status
	}
	fmt.Printf("Product ID #%d is %s\n", pid, status)

	// for case #3
	//     What is Greco's second phone number?
	name, index, phoeno := "greco", 2, "N/A"
	if v, ok := mulphones[name]; ok {
		if index < len(v) {
			phoeno = v[index]
		}
	}
	fmt.Printf("%s's %dnd phone number: %s\n", name, index, phoeno)

	// for case #4
	//     How many of 576872813 the customer 101 is going to buy?
	//                (Product ID)  (Customer ID)

	cid := 101
	prid := 576872813
	var quantity int
	if v, ok := basket[cid]; ok {
		if q, ok := v[prid]; ok {
			quantity = q
		}
	}

	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, quantity, prid)
}
