package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  int64
}

func (c *SafeCounter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v++
}

type AtomicCounter struct {
	v int64
}

func (c *AtomicCounter) Increase() {
	atomic.AddInt64(&c.v, 1)
}

func main() {
	counter := SafeCounter{v: 0}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				counter.Increase()
			}
		}()
	}

	atomicCounter := AtomicCounter{v: 0}
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				atomicCounter.Increase()
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("SafeCounter最终结果：", counter.v)
	fmt.Println("AtomicCounter最终结果：", atomicCounter.v)
}
