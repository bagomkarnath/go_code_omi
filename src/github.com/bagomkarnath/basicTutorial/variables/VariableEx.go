package main

import "fmt"

// package level variable
var test int = 50

// var grouping
var (
	name string = "Scott"
	age  int    = 58
	city string = "Vizag"
)

func main() {
	var i int
	i = 40
	fmt.Printf("%v %T \n", i, i)

	var j int = 20
	fmt.Printf("%v %T \n", j, j)

	k := 60
	fmt.Printf("%v %T \n", k, k)

	var l float32 = 20
	fmt.Printf("%v %T \n", l, l)

	m := 60.
	fmt.Printf("%v %T \n", m, m)

	fmt.Printf("%v %T \n", test, test)
}
