package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("slice copying....")

	evens := []int{2, 4, 6}
	odds := []int{3, 5, 7}

	s.PrintBacking = true
	s.Show("evens {before}", evens)
	s.Show("odds {before}", odds)

	// copy odds to evens
	N := copy(evens, odds) // copy(target, source)
	fmt.Printf("%d elements copied from odds to evens\n", N)

	s.Show("evens {after}", evens)
	s.Show("odds {after}", odds)

	// raining probabilities
	data := []float64{10, 25, 30, 50}
	// s.Show("data {before}", data)

	// received new data
	// newData := []float64{80, 90}
	// for i := 0; i < len(newData); i++ {
	// 	data[i] = newData[i]
	// }

	// using copy instead
	// n := copy(data, []float64{10, 5, 15, 0, 20})
	// fmt.Printf("%d elements copied\n", n)

	// using append instead of copy
	// data = append(data, []float64{10, 5, 15, 0, 20}...)
	// data = append(data[:0], []float64{10, 5, 15, 0, 20}...)

	// creating clone of slice
	// saved := make([]float64, len(data))
	// copy(saved, data)

	// clone a slice using append nil
	// saved := append([]float64(nil), data...)
	// s.Show("data {after}", saved)

	s.Show("Probabilities", data)
	fmt.Printf("Is it going to rain? %.f%% chance\n",
		(data[0]+data[1]+data[2]+data[3])/float64(len(data)))

	lyric := []string{"show", "me", "my", "silver", "lining"}
	s.Show("lyric", lyric)
	part := lyric[:2:2]
	fmt.Println(part)
	s.Show("part", part)
	part = append(part, "right", "place")
	s.Show("part", part)

	tasks1 := make([]string, 2)
	s.Show("tasks1", tasks1)
	tasks1 = append(tasks1, "hello", "world")
	s.Show("tasks1", tasks1)
	fmt.Printf("%q\n", tasks1)

	tasks2 := make([]string, 0, 2)
	s.Show("tasks", tasks2)
	tasks2 = append(tasks2, "hello", "world")
	s.Show("tasks", tasks2)
	fmt.Printf("%q\n", tasks2)

	song := []string{"le", "vent", "nous", "portera"}
	s.Show("song", song)
	n := copy(song, make([]string, 4))
	s.Show("song", song)

	fmt.Printf("%d %q\n", n, song)

	//
	spendings := [][]int{{1, 2}}
	s.Show("spendings", spendings)
	// REMEMBER: %T prints the type of a given value
	fmt.Printf("%T - ", spendings)
	fmt.Printf("%T - ", spendings[0])
	fmt.Printf("%T", spendings[0][0])

}
