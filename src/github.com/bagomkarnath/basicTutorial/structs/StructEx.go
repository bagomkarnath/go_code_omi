package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Dr. Vijay",
		companions: []string{
			"Madhu", "Jadu", "Laddu",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	//anonymStruct := struct{ name string }{name: "Dr Vijay"}
	//fmt.Println(anonymStruct)

}
