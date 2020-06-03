package main

import "fmt"

func main() {
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)

	(*ms).foo = 43         //ms.foo will also work ... compiler helps
	fmt.Println((*ms).foo) //ms.foo
}

type myStruct struct {
	foo int
}
