package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicer()
	fmt.Println("End")
}

func panicer() {
	fmt.Println("panic is about to start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error : ", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicing")
}
