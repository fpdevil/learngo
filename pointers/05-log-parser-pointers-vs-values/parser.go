package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string // the unique domain name from log
	visits int    // number of visits for the domain
	// add more metrics if needed later
}

// parser keep tracks of the log file parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
	lerr    error             // the last error occurred
}

// newParser function creates constructor and initializes the
// parser struct which can be populated with data
func newParser() *parser {
	return &parser{sum: make(map[string]result)}
}

// parseLines function takes the parer instance and an input string
// read from the log file and checks as well as parses the file
func parseLines(p *parser, input string) (r result) {
	// error handling
	if p.lerr != nil {
		return
	}

	// update the line counter for each line read
	p.lines++

	// split the input fields from the log file line by line
	fields := strings.Fields(input)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	var err error
	r.domain = fields[0]
	// sum the total visits per domain
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}

	return
}

// update the results from the parser and return the new updated
// parser populated with data from log file
func update(p *parser, r result) {

	if p.lerr != nil {
		return
	}

	// Collect the unique domains
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// Keep track of total and per domain visits
	p.total += r.visits

	// You cannot assign to composite values
	// p.sum[domain].visits += visits

	// create and assign a new copy of `visit`
	res := result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
	p.sum[r.domain] = res
}

// summarize function summarizes and prints the results
func summarize(p *parser) {
	// Print the visits per domain
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

// dumpErrs is the helper function which simplifies handling
// multiple errors
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Print("> Err:", err)
		}
	}
}

// err returns the last error encountered
func err(p *parser) error {
	return p.lerr
}
