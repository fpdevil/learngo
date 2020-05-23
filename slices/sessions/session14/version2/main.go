package main

import "fmt"

func main() {
	// spending for 5 days
	// data from a source
	spendings := make([][]int, 0, 5)

	spendings = append(spendings, []int{200, 100})
	spendings = append(spendings, []int{25, 10, 45, 60})
	spendings = append(spendings, []int{5, 15, 35})
	spendings = append(spendings, []int{95, 10}, []int{50, 25})

	for i, spent := range spendings {
		var total int
		for j := range spent {
			total += spent[j]
		}
		fmt.Printf("Day %d: %d\n", i, total)
	}
}
