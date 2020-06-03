package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0

func main() {
	fmt.Println("Start")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello : ", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
