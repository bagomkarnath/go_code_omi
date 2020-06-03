package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0
var m = sync.RWMutex{}

func main() {
	fmt.Println("Start")
	//runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello : ", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
