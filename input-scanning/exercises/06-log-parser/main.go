package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	fmt.Println("--- Exerises 23-input-scanning/06-log-parser ---")
	fmt.Println("\tLog Parser from Stratch")
	fmt.Printf("%s\n", strings.Repeat("=", 45))
	fmt.Println("\n")

	var (
		// a map for holding the domain and number of visits
		logger map[string]int

		// for counting the lines
		count int

		// domains slice will contain domains list from the file
		domains []string

		// total holds the sum of all visits
		total int
	)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	// initialize the map
	logger = make(map[string]int)

	for in.Scan() {
		count++

		// split each line into domain and visits
		fields := strings.Fields(in.Text())
		// fmt.Printf("%#v\n", fields)

		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, count)
			return
		}

		domain, visitstr := fields[0], fields[1]
		visits, err := strconv.Atoi(visitstr)

		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", visitstr, count)
			return
		}

		// insert into the domain slice only if the domain is
		// not yet present in the map to avoid duplicates
		if _, ok := logger[domain]; !ok {
			domains = append(domains, domain)
		}

		logger[domain] += visits
	}

	// handle input errors
	if err := in.Err(); err != nil {
		fmt.Printf("input error: %s\n", err)
		return
	}

	// sort the domains
	sort.Strings(domains)
	// fmt.Printf("domains: %#v\n", domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Printf("%s\n", strings.Repeat("-", 45))

	for _, domain := range domains {
		visits := logger[domain]
		total += visits
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Println()
	fmt.Printf("%-30s %10d\n", "TOTAL", total)

	// for domain, visits := range logger {
	// 	fmt.Printf("%-30s %10d\n", domain, visits)
	// }
}
