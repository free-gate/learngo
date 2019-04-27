// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		update(p, parse(p, in.Text()))
	}

	summarize(p)
	dumpErrs([]error{in.Err(), err(p)})
}

// summarize summarizes and prints the parsing result
func summarize(p *parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

// dumpErrs simplifies handling multiple errors
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}