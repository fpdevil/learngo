package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("Handling slice expressions...")
	s.PrintBacking = true

	nums := []int{1, 3, 2, 4}
	s.Show("nums", nums)

	// odds := nums[:2]
	// s.Show("odds", odds)
	// odds                (len:2  cap:4  ptr:9776)
	// ╔═══╗╔═══╗+---++---+
	// ║ 1 ║║ 3 ║| 2 || 4 |
	// ╚═══╝╚═══╝+---++---+
	//   0    1    2    3

	// using slicing expressions
	odds := nums[:2:2]

	// append 2 more elements to the odds
	odds = append(odds, 5, 7)

	// append 2 for evens
	evens := append(nums[2:4], 6, 8) // this is not full slice expression

	s.Show("odds", odds)
	s.Show("nums", nums)
	s.Show("evens", evens)
	s.Show("nums", nums)

	// make function usage
	tasks := []string{"eat", "sleep", "jump"}
	s.Show("tasks", tasks)

	// convert tasks to upper
	// upTasks := make([]string, len(tasks))
	upTasks := make([]string, 0, cap(tasks))

	// increase capacity of upTasks
	// upTasks = upTasks[:cap(upTasks)]
	s.Show("upTasks initial", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		// upTasks[i] = strings.ToUpper(task)
		s.Show("upTasks", upTasks)
	}
}
