package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		id := rnd.Intn(10) + 1
		// sync: negative WaitGroup counter
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				cache[id] = b
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		// fatal error: concurrent map writes
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}

	return Book{}, false
}
