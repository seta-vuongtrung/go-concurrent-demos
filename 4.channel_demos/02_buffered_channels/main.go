package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27
		// fatal error: all goroutines are asleep - deadlock!
		// ch <- 12
		// ch <- 24
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
