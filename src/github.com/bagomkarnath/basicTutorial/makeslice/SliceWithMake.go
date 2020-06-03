package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("length : %v \n", len(a))
	fmt.Printf("capacity : %v \n", cap(a))

	b := []int{}
	fmt.Println(b)
	fmt.Printf("length : %v \n", len(b))
	fmt.Printf("capacity : %v \n", cap(b))
	b = append(b, 1)
	fmt.Println(b)
	fmt.Printf("length : %v \n", len(b))
	fmt.Printf("capacity : %v \n", cap(b))
	b = append(b, 2, 4, 5, 67, 78)
	fmt.Println(b)
	fmt.Printf("length : %v \n", len(b))
	fmt.Printf("capacity : %v \n", cap(b))

	//a := append(a, []int{1,2,3}) //error
	a = append(a, []int{1, 2, 3}...)
	fmt.Println(a)
}
