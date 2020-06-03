package main

import (
	"fmt"
)

const (
	a = iota + 5
	b
	c
)

func main() {
	fmt.Printf("%v , %v, %v", a, b, c)
}
