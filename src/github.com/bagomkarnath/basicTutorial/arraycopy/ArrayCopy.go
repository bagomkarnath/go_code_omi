package main

import (
	"fmt"
)

func main() {
	nameArr := [...]string{"ramesh", "suresh", "ganesh"}
	copiedArr := nameArr //copies value and creates a new array
	fmt.Printf("nameArr %v \n", nameArr)
	fmt.Printf("copiedArr %v \n", copiedArr)
	// update value in copied array
	copiedArr[1] = "SACHIN"
	fmt.Printf("nameArr %v \n", nameArr)
	fmt.Printf("copiedArr %v \n", copiedArr)
	//using pointer
	copiedArr2 := &nameArr
	copiedArr2[1] = "Ganguly"
	fmt.Printf("nameArr %v \n", nameArr)
	fmt.Printf("copiedArr %v \n", copiedArr2)
}
