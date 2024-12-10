package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"weightconversion/weightconv"
)

func main() {
	data := os.Args[1:]
	if len(data) > 0 {
		err := outputInfo(strings.NewReader(strings.Join(data, "\n")))
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	} else {
		err := outputInfo(os.Stdin)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
	}
}

func outputInfo(data io.Reader) error {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		w, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Print(err)
			os.Exit(1)
		}
		k := weightconv.Kilogram(w)
		p := weightconv.Pound(w)
		fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
	}
	return nil
}
