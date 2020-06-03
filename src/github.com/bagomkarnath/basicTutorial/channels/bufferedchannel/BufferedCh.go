package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) { //Recieve only channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { //Send only channel
		ch <- 42
		ch <- 28
		ch <- 100
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
