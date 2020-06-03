package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 32
	var j float32
	j = float32(i)

	fmt.Printf("%v %T  \n", i, i)
	fmt.Printf("%v %T  \n", j, j)

	var k string

	k = string(i)

	fmt.Printf("%v %T  \n", k, k)
	k = strconv.Itoa(i)

	fmt.Printf("%v %T  \n", k, k)
}
