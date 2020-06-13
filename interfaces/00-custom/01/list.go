package main

import "fmt"

type list []Speaker

func (l list) speaks() {
	for _, v := range l {
		fmt.Printf("%-15T  >> ", v)
		v.Speak()
	}
}
