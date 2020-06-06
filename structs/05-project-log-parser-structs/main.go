package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// result struct for storing the visits for each domain
type result struct {
	domain string // domain from the log file
	visits int    // number of visits per domain
}

// parser - a struct for keeping details of log file parsing
type parser struct {
	sum     map[string]result // holds the details of metrics per domain
	domains []string          // unique domain names
	total   int               // total visits from all domains
	lines   int               // number of parsed log file entries
}

func main() {
	fmt.Println("24-structs/05-project-log-parser-structs")

	p := parser{
		// initialize the sum map declared
		sum: make(map[string]result),
	}

	// create a new scanner for input data handling
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	// scan input data line by line
	for in.Scan() {
		// increment the line count
		p.lines++

		// split the line elements
		fields := strings.Fields(in.Text())

		// check to ensure that there are 2 fields per each line
		if len(fields) != 2 {
			fmt.Printf("ERROR: Wrong Input %v (line: %#v)\n", fields, p.lines)
			return
		}

		// data conversion and handline
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])

		if visits < 0 || err != nil {
			fmt.Printf("ERROR: Invalid Input %q (line: %#v)\n", fields[1], p.lines)
			return
		}

		// now append to the slice only unique domain names
		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}

		// aggregate the total viists for all domains
		p.total += visits
		// sum[domain] += visits

		// we cannot assign to composite values
		// p.sum[domain].visits += visits

		// create a cpoy of the results
		r := result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
		p.sum[domain] = r
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
