package main

import (
	"fmt"
	"strconv"
	"strings"
)

// result struct for storing the visits for each domain
type result struct {
	domain string // domain string from the log file
	visits int    // number of visits per domain from log file
}

// parser - a struct for keeping details of log file parsing
type parser struct {
	sum     map[string]result // holds the details of metrics per domain
	domains []string          // unique domain names
	total   int               // total visits from all domains
	lines   int               // number of parsed log file entries
}

// newParser function bootstraps the parser initializing the map
func newParser() parser {
	p := parser{
		// initialize the sum map declared
		sum: make(map[string]result),
	}
	return p
}

// parse parses a single log line and returns the result with an error value
func parse(p parser, line string) (parsed result, err error) {

	// inside the function, the result values parsed and
	// err become normal variables to work with
	// we declare these as return values now
	// var (
	// 	parsed result
	// 	err    error
	// )

	// split the line elements
	fields := strings.Fields(line)

	// check to ensure that there are 2 fields per each line
	if len(fields) != 2 {
		err = fmt.Errorf("Wrong Input %v (line: %#v)", fields, p.lines)
		// we dont have to specify the return values as they were
		// defined globally at function definition level now
		// return parsed, err
		return
	}

	// data conversion and handline
	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])

	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("Invalid Input %q (line: %#v)", fields[1], p.lines)
		// we dont have to specify the return values as they were
		// defined globally at function definition level now
		// return parsed, err
		return
	}

	// return result{domain: parsed.domain, visits: parsed.visits}
	// we dont have to specify the return values as they were
	// defined globally at function definition level now
	// return parsed, err
	return
}

// update function handles updating of the results for the
// parser result
func update(p parser, parsed result) parser {
	domain, visits := parsed.domain, parsed.visits

	// now append to the slice only the unique domain names
	// we will first if the sum map already has the domain or
	// not and then append the value
	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	// now aggregate the total number of viists for all domains
	// giving the total number of visits overall
	p.total += visits

	// create a cpoy of the results
	r := result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
	p.sum[domain] = r

	return p
}
