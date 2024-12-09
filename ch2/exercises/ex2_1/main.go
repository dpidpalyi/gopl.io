package main

import (
	"fmt"

	"tempconversion/tempconv"
)

func main() {
	c := tempconv.FToC(212.0)
	fmt.Println(c)
	fmt.Println(tempconv.CToK(c))
}
