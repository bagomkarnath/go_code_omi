package main

import (
	"fmt"
)

func main() {
	var names [5]string

	names[0] = "Ramesh"
	names[1] = "Suresh"
	names[2] = "Ganesh"

	fmt.Printf("%v %T \n", names, names)

	grades := [3]string{"A", "A", "C"}
	fmt.Printf("%v %T \n", grades, grades)

	marks := [...]int{59, 74, 87}
	fmt.Printf("%v %T \n", marks, marks)

	fmt.Printf("length of names array %v \n", len(names))
	fmt.Printf("length of grades array %v \n", len(grades))
	fmt.Printf("length of marks array %v \n", len(marks))
}
