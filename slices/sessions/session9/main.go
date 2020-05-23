package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("Going through append...")

	s.PrintBacking = true
	s.MaxPerLine = 4
	// s.PrintElementAddr = true

	var items []string
	s.Show("items nil slice", items)

	items = append(items, "toys")
	s.Show("items now", items)

	items = append(items, "puzzles", "boards")
	s.Show("items after adding 2", items)

	groups := []string{"indoor", "outdoor"}
	items = append(items, groups...)

	s.Show("items", items)

	// adding items in the middle
	var nums []int
	s.Show("No Backing Array yet for nums", nums)

	nums = append(nums, 1, 3)
	s.Show("nums after 1st append", nums)

	nums = append(nums, 2)
	s.Show("append 2", nums)

	nums = append(nums, 4)
	s.Show("append 4", nums)

	// now add 10 and 11 in 2nd and 3rd places
	// But this overrides the existing values 2 and 4
	// instead of inserting new values
	// No Backing Array yet for nums  (len:0  cap:0  ptr:0   )
	// <nil slice>
	//  nums after 1st append  (len:2  cap:2  ptr:9536)
	// ╔═══╗╔═══╗
	// ║ 1 ║║ 3 ║
	// ╚═══╝╚═══╝
	//   0    1
	//  append 2            (len:3  cap:4  ptr:1792)
	// ╔═══╗╔═══╗╔═══╗+---+
	// ║ 1 ║║ 3 ║║ 2 ║| 0 |
	// ╚═══╝╚═══╝╚═══╝+---+
	//   0    1    2    3
	//  append 4            (len:4  cap:4  ptr:1792)
	// ╔═══╗╔═══╗╔═══╗╔═══╗
	// ║ 1 ║║ 3 ║║ 2 ║║ 4 ║
	// ╚═══╝╚═══╝╚═══╝╚═══╝
	//   0    1    2    3
	//  nums[:2] <= 10, 11  (len:4  cap:4  ptr:1792)
	// ╔═══╗╔═══╗╔════╗╔════╗
	// ║ 1 ║║ 3 ║║ 10 ║║ 11 ║
	// ╚═══╝╚═══╝╚════╝╚════╝
	//   0    1     2     3
	// nums = append(nums[:2], 10, 11)
	// s.Show("nums[:2] <= 10, 11", nums)

	// first append the items after position 2 after the nums
	// by duplicating them first then we can override their
	// existing places with new values as under.
	nums = append(nums, nums[2:]...)
	s.Show("Duplicating nums <= nums[2:]", nums)

	// now insert the elements
	nums = append(nums[:2], 10, 11)
	s.Show("nums[:2] <= 10, 11", nums)

	// This shows the below map
	// nums[:2] <= 10, 11  (len:4  cap:8  ptr:4224)
	// ╔═══╗╔═══╗╔════╗╔════╗+---+
	// ║ 1 ║║ 3 ║║ 10 ║║ 11 ║| 2 |
	// ╚═══╝╚═══╝╚════╝╚════╝+---+
	//   0    1     2     3    4
	// +---++---++---+
	// | 4 || 0 || 0 |
	// +---++---++---+
	//   5    6    7

	// for revealing the hidden elements, extend the slice as below
	nums = nums[:6]
	s.Show("slice extended nums[:6]", nums)

	words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
	s.Show("pre append words", words)

	fmt.Println("words[:1] =", words[:1])
	words = append(words[:1], "is", "everywhere")
	s.Show("post append words", words)

	fmt.Println("words[len(words) + 1 : cap(words)] =", words[len(words)+1:cap(words)])
	words = append(words, words[len(words)+1:cap(words)]...)
	s.Show("words", words)

	// keyed slices
	junks := []string{9: ""}
	fmt.Printf("length: %d, capacity: %d\n", len(junks), cap(junks))
	s.Show("keyed slice junks", junks)

	junks = append(junks, "boom!")
	fmt.Printf("length: %d, capacity: %d\n", len(junks), cap(junks))
	s.Show("junks now", junks)
}
