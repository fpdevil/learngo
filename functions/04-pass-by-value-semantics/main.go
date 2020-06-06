package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("--- 04-pass-by-value-semantics ---")
	fmt.Println("	LOG PARSER REVISTED	")
	fmt.Println()

	// create a new parser instance
	p := newParser()

	// create a new scanner for input data handling
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	// scan input data line by line
	for in.Scan() {
		// increment the lines count to track which line under context
		p.lines++

		parsed, err := parse(p, in.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		p = update(p, parsed)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Printf("%s\n", strings.Repeat("-", 45))

	// sort the domain names
	sort.Strings(p.domains)

	// now loop over the slice for each domain name
	for _, domain := range p.domains {
		v := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, v.visits)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-30s %10d\n", "TOTAL", p.total)

	// error handling for the scanner
	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %q\n", err)
	}
}
