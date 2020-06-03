package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

// Defer will execute after main but before it returns anything
