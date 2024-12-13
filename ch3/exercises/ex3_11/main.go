// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign byte
	n := len(s)
	if n > 0 && (s[0] == '-' || s[0] == '+') {
		sign = s[0]
		s = s[1:]
	}

	var buf bytes.Buffer
	buf.WriteByte(sign)

	if period := strings.LastIndex(s, "."); period != -1 {
		writeComma(s[:period], &buf)
		buf.WriteByte('.')
		writeComma(s[period+1:], &buf)
	} else {
		writeComma(s, &buf)
	}
	return buf.String()
}

func writeComma(s string, buf *bytes.Buffer) {
	d, t := len(s)%3, len(s)/3
	if d != 0 {
		buf.WriteString(s[:d])
		buf.WriteByte(',')
	}

	for i := 0; i < t; i++ {
		buf.WriteString(s[i*3+d : d+i*3+3])
		if i != t-1 {
			buf.WriteByte(',')
		}
	}
}

//!-
