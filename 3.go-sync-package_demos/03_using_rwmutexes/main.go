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
	m := &sync.RWMutex{}
	for i := 0; i < 20; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(wg, m)
		go func(wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				m.Lock()
				cache[id] = b
				m.Unlock()
				fmt.Println(b)
			}
			wg.Done()
		}(wg, m)
		if i == 10 {
			time.Sleep(150 * time.Millisecond)
		}
	}

	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
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
