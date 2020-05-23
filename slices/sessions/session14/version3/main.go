package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Spendings revisited...")

	spendings := fetch()

	for i, spent := range spendings {
		var total int
		for j := range spent {
			total += spent[j]
		}
		fmt.Printf("Day: %d: %d\n", i, total)
	}
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))
	// fmt.Printf("%#v\n", spendings)

	for i, line := range lines {
		// fmt.Printf("%d: %#v\n", i, line)

		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			// fmt.Printf("\t %d: %#v\n", j, field)

			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}

	// return nil
	return spendings
}
