package main

import "fmt"

func main() {
	fmt.Println("--- 26-pointers/02-pointers-basic-examples ---")
	fmt.Println()

	// checking of pointer is nil (zero value)
	// with P as *int
	// var P *int
	// P is nil and it's address is 0x0
	//

	var (
		counter int
		P       *int
	)

	counter = 100
	P = &counter // P is a pointer to int variable
	V := *P      // V is an int variable (a copy of counter)

	// checking of pointer is nil (zero value)
	// with P as *int
	// var P *int
	// P is nil and it's address is 0x0
	//
	if P == nil {
		fmt.Printf("P is nil and it's address is %p\n", P)
	}

	if P == &counter {
		fmt.Printf("P is equal to counter's address: %p = %p\n", P, &counter)
	}

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	counter = 10
	fmt.Println()
	fmt.Println("*** change counter ***")
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println()
	fmt.Println("*** change counter in passVal() ***")
	passVal(counter)
	fmt.Println()
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println()
	fmt.Println("*** change counter assigning return value from passVal() ***")
	counter = passVal(counter)
	fmt.Println()
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)

	fmt.Println()
	fmt.Println("*** change counter sending address to passPtrVal() ***")
	passPtrVal(&counter)
	passPtrVal(&counter)
	passPtrVal(&counter)
	passPtrVal(&counter)
	passPtrVal(&counter)
	fmt.Println()
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)
}

// function to pass by value (default)
// n is a int variable (copy of counter)
func passVal(n int) int {
	// when this is called, the value of counter doesn't
	// change because `n` is a copy
	n = 50
	fmt.Printf("n      : %-13d addr: %-13p\n", n, &n)
	return n
}

// for sending a pointer to a function
// *pn is a int pointer variable (copy of P)
func passPtrVal(pn *int) {
	// fmt.Printf("pn     : %-13p addr: %-13p *pn: %-13d\n", pn, &pn, *pn)

	// using the indirect operation `*pn` update counter
	// the value of counter changes because the pointer
	// `pn` points to `counter` — (*pn)++
	*pn++ // (*pn)++    ===> go converts to *pn++

	fmt.Printf("pn     : %-13p addr: %-13p *pn: %-13d\n", pn, &pn, *pn)
	// setting pn to nil won't affect counter as it just
	// updates the local variable
	pn = nil
}
