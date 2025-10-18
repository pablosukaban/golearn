package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	//counter := SafeCounter{v: make(map[string]int)}
	//
	//for i := 0; i < 1000; i++ {
	//	go counter.Inc("some key")
	//}
	//
	//time.Sleep(time.Second)
	//
	//fmt.Println(counter.Value("some key"))

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
		go c.Inc("other")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
	fmt.Println(c.Value("other"))
}
