package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n, err")
				continue
			}
			if isDupExist := countLines(f, counts); isDupExist {
				fmt.Printf("duplicate lines exist in: %v\n", arg)
			}
		}
	}
	for line, cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%v\t%v\n", cnt, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool {
	isDupExist := false
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !isDupExist && counts[line] > 1 {
			isDupExist = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	return isDupExist
}
