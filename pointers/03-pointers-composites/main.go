package main

import (
	"fmt"
	"strings"
)

type house struct {
	name  string
	rooms int
}

func main() {
	fmt.Println("---- 26-pointers/03-pointers-composites ----")
	fmt.Println()

	fmt.Println("••••• ARRAYS •••••")
	arrays()
	fmt.Println()

	fmt.Println("••••• SLICES •••••")
	slices()
	fmt.Println()

	fmt.Println("••••• MAPS •••••")
	maps()
	fmt.Println()

	fmt.Println("••••• STRUCTS •••••")
	structs()
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)
	fmt.Printf("structs()       : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("structs()       : %p %+v\n", &myHouse, myHouse)
}

func addRoom(h house) {
	h.rooms++
	fmt.Printf("addRoom()       : %p %+v\n", &h, h)
}

func addRoomPtr(h *house) {
	(*h).rooms++ // both are same as: (*h).rooms++ and h.rooms++
	fmt.Printf("addRoomPtr()    : %p %+v\n", h, h)
	fmt.Printf("&h.name         : %p\n", &h.name)
	fmt.Printf("&h.rooms        : %p\n", &h.rooms)
}

func maps() {
	confused := map[string]int{"one": 2, "two": 0}
	fmt.Printf("maps            : %p %#v\n", &confused, confused)

	fix(confused)
	fmt.Printf("maps            : %p %#v\n", &confused, confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Printf("fix(maps)       : %p %#v\n", &m, m)
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	fmt.Printf("slices list     : %p %q\n", &dirs, dirs)

	upPtr(&dirs)
	fmt.Printf("slices list     : %p %q\n", &dirs, dirs)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}

	// appending a value to the slice
	list = append(list, "HEISEN BUG")

	fmt.Printf("up(dirs)        : %p %q\n", &list, list)
}

func upPtr(list *[]string) {
	lv := *list // create a pointer to list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	// appending a value to the slice
	// *list = append(*list, "HEISEN BUG")
	*list = append(*list, "HEISEN BUG")

	fmt.Printf("upPtr(&dirs)    : %p %q\n", list, list)
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("arrays nums     : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("incr nums       : %p\n", &nums)
	for i := range nums {
		nums[i]++
		fmt.Printf("incr.nums[%d]    : %p\n", i, &nums[i])
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums  : %p\n", &nums)
	for i := range nums {
		// nums[i]++ is same as: (*nums)[i]++
		(*nums)[i]++
	}
}
