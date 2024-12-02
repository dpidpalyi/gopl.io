package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(sl []string) string {
	var s, sep string
	for i := 0; i < len(sl); i++ {
		s += sep + sl[i]
		sep = " "
	}
	return s
}

func echo2(sl []string) string {
	var s, sep string
	for _, arg := range sl {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3(sl []string) string {
	return strings.Join(sl, " ")
}

func main() {
	start := time.Now()
	fmt.Println(echo1(os.Args))
	fmt.Printf("%vms elapsed\n", time.Since(start).Nanoseconds())
	start = time.Now()
	fmt.Println(echo2(os.Args))
	fmt.Printf("%vms elapsed\n", time.Since(start).Nanoseconds())
	start = time.Now()
	fmt.Println(echo3(os.Args))
	fmt.Printf("%vms elapsed\n", time.Since(start).Nanoseconds())
}
