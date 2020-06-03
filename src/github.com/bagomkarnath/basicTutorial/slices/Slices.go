package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 4}
	fmt.Printf("%v %T\n", slice1, slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// slices are by default reference type

	sliceCopy := slice1

	sliceCopy[1] = 67

	fmt.Println(sliceCopy)
	fmt.Println(slice1)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
