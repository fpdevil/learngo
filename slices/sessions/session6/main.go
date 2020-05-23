package main

import (
	"fmt"
	"sort"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("Slice Internals: Backing Array....")

	// declare an arbitrary slice
	// grades := [...]float64{40, 10, 20, 50, 60, 70}
	grades := []float64{40, 10, 20, 50, 60, 70}

	// creating a new nil slice this will break the connection
	// with the backing array so that changing this does
	// affect the original
	var newGrades []float64
	newGrades = append(newGrades, grades...)

	// another shorter way of achieving the same
	// newGrades := append([]float64(nil), grades...)

	// front := grades[:3] // ---> this points to the original grades
	front := newGrades[:3] // ---> poinitng to the new grades
	front2 := front[:3]
	front3 := front

	// sort the front
	sort.Float64s(front)

	s.PrintBacking = true
	s.MaxPerLine = 7
	s.Show("grades", grades[:])
	s.Show("newGrades", newGrades)
	s.Show("front", front)
	s.Show("front2", front2)
	s.Show("front3", front3)

	arr := [...]int{9, 7, 5, 3, 1}
	s.Show("arr", arr)
	nums := arr[2:]
	s.Show("nums", nums)
	nums2 := nums[1:]
	s.Show("nums2", nums2)

	arr[2]++
	s.Show("arr", arr)
	nums[1] -= arr[4] - 4
	s.Show("nums", nums)
	nums2[1] += 5
	s.Show("nums2", nums2)
	s.Show("arr", arr)
	fmt.Println(nums)
}
