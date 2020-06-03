package main

import "fmt"

func main() {
	var a int = 47
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 90
	fmt.Println(a, *b)

}
