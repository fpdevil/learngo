package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parsing a log file and generating a report
// NOTE: Never loop over a map [a serious design mistake]

func main() {
	fmt.Println("---- 23-input-scanning/03-project-log-parser ----")
	fmt.Println()

	// create a map for holding the unique domain name and visits
	var (
		sum map[string]int

		// create a slice to store unique domain names
		domains []string

		// aggregate of total visits for each and every domain
		total int

		// capture line number of error
		errline int
	)
	// initialize the map
	sum = make(map[string]int)

	// create a new scanner for input data
	in := bufio.NewScanner(os.Stdin)

	// scan the file line-by-line
	for in.Scan() {
		errline++

		fields := strings.Fields(in.Text())

		// check for the number of fields
		if len(fields) != 2 {
			fmt.Printf("ERROR: Wrong Input %v (line: %#v)\n", fields, errline)
			return
		}


		// fmt.Printf("domain: %s(%[1]T) - visits: %s(%[2]T)\n", fields[0], fields[1])

		// data conversion and handling
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			// fmt.Printf("ERROR: %s\n", err)
			fmt.Printf("ERROR: Invalid Input %q (line: %#v)\n", fields[1], errline)
			return
		}

		// append to the slice unique domain names
		if _, ok := sum[domain]; ok {
			domains = append(domains, domain)
		}

		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Printf("%s\n", strings.Repeat("-", 45))

	// with map, the results irder change each time
	// not a god option to iterate over the map
	// for domain, visit := range sum {
	// 	fmt.Printf("%-30s %10d\n", domain, visit)
	// }

	// sort the domain names
	sort.Strings(domains)

	// now loop over the slice for each domain name
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-30s %10d\n", "TOTAL", total)

	// error handling for the scanner
	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %q\n", err)
	}
}
